package easyweb

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/lengzhao/easyweb/util"
)

var _ Page = &easyPage{}

func newPage(conn *websocket.Conn) *easyPage {
	var page easyPage
	page.conn = conn
	page.callback = make(map[string]eventMsgData)
	page.msgChan = make(chan any, 10)
	page.closed = make(chan int)
	page.env = make(map[string]any)
	page.watchEnv = make(map[string]func(value any))
	return &page
}

func (p *easyPage) Title(title string) Page {
	id := util.GetCallerID(util.LevelParent)
	msg := toClientMsgData{id, "title", title}
	p.sendMsg(msg)
	return p
}

func (p *easyPage) AddJs(js string) Page {
	id := util.GetCallerID(util.LevelParent)
	msg := toClientMsgData{id, "js", js}
	p.sendMsg(msg)
	return p
}

func (p *easyPage) RunJs(js string) Page {
	id := util.GetCallerID(util.LevelParent)
	msg := toClientMsgData{id, "eval", js}
	p.sendMsg(msg)
	return p
}

func (p *easyPage) GetPeer() string {
	return p.conn.RemoteAddr().String()
}

func (p *easyPage) AddCss(css string) Page {
	id := util.GetCallerID(util.LevelParent)
	msg := toClientMsgData{id, "css", css}
	p.sendMsg(msg)
	return p
}

// e 将作为子元素放到div内，如<div id=rand_id>e</div>
func (p *easyPage) Write(e any) string {
	var id string
	gid, ok := e.(IGetID)
	if ok {
		id = gid.GetID() + "_p"
	} else {
		id = util.GetCallerID(util.LevelParent)
	}
	return p.WriteWithID(id, e)
}

// e 将作为子元素放到id所属的元素内，如<div id=id>e</div>
func (p *easyPage) WriteWithID(id string, e any) string {
	msg := toClientMsgData{id, "", fmt.Sprint(e)}
	p.sendMsg(msg)
	if e, ok := e.(IAfterLoaded); ok {
		e.AfterElementLoadedFromFramwork(p)
	}
	return id
}

type attrInfo struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// 修改指定id的元素的属性
func (p *easyPage) SetAttr(id, key, value string) string {
	info := attrInfo{Key: key, Value: value}
	data, _ := json.Marshal(info)
	msg := toClientMsgData{id, "attr", string(data)}
	p.sendMsg(msg)
	return id
}

// 注册指定id的前端事件，typ是前端事件名，cb是回调还是，如果cb为空，则是取消事件回调
func (p *easyPage) RegistEvent(id, typ string, cb IMessageCb) {
	if id == "" || typ == "" {
		return
	}
	// 1. add event callback(server side)
	cbMsg := eventMsgData{}
	cbMsg.ID = id
	cbMsg.Event = cb
	p.sendMsg(cbMsg)
	// 2. add client event(jquery will handle the event)
	if cb != nil {
		toClient := toClientMsgData{id, "event", typ}
		p.sendMsg(toClient)
	} else {
		toClient := toClientMsgData{id, "off", typ}
		p.sendMsg(toClient)
	}
}

func (p *easyPage) sendMsg(msg any) {
	select {
	case <-p.closed:
	case p.msgChan <- msg:
	}
}

func (p *easyPage) Close() {
	select {
	case <-p.closed:
	default:
		p.conn.Close()
	}
}

func (p *easyPage) WaitUntilClosed() {
	<-p.closed
}

func (p *easyPage) SetEnv(key string, value any) {
	p.mu.Lock()
	p.env[key] = value
	cb := p.watchEnv[key]
	p.mu.Unlock()
	if cb != nil {
		cb(value)
	}
}

func (p *easyPage) GetEnv(key string) (value any) {
	p.mu.Lock()
	value = p.env[key]
	p.mu.Unlock()
	return
}

func (p *easyPage) WatchEnv(key string, cb func(value any)) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.watchEnv[key] != nil {
		return fmt.Errorf("exist callback")
	}
	p.watchEnv[key] = cb
	return nil
}

func encode(v interface{}) []byte {
	buff := new(strings.Builder)
	enc := json.NewEncoder(buff)
	enc.SetEscapeHTML(false)
	enc.Encode(v)
	return []byte(buff.String())
}

func (p *easyPage) processMsg() {
	for {
		select {
		case <-p.closed:
			return
		case data, ok := <-p.msgChan:
			if !ok {
				continue
			}
			switch msg := data.(type) {
			case toClientMsgData:
				if msg.ID == "" {
					continue
				}
				p.conn.WriteMessage(websocket.TextMessage, encode(msg))
				if msg.Msg == "" {
					delete(p.callback, msg.ID)
				}
			case fromClientMsgData:
				cb := p.callback[msg.ID]
				if cb.Event != nil {
					go func(id string, dataType CbDataType, data []byte) {
						defer func() {
							if r := recover(); r != nil {
								fmt.Println("Recovered", r)
							}
						}()
						cb.Event.MessageCallbackFromFramwork(msg.ID, dataType, data)
					}(msg.ID, CbDataTypeString, []byte(msg.Msg))
				}
			case fileMsgData:
				cb := p.callback[msg.ID]
				if cb.Event != nil {
					go func(id string, dataType CbDataType, data []byte) {
						defer func() {
							if r := recover(); r != nil {
								fmt.Println("Recovered", r)
							}
						}()
						cb.Event.MessageCallbackFromFramwork(msg.ID, dataType, data)
					}(msg.ID, CbDataTypeBinary, msg.BinaryData)
				}
			case eventMsgData:
				if msg.Event == nil {
					delete(p.callback, msg.ID)
				} else {
					p.callback[msg.ID] = msg
				}
			default:
				fmt.Println("unknown msg type:", msg)
			}
		}
	}
}

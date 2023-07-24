package easyweb

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/lengzhao/easyweb/util"
)

var _ Page = &easyPage{}

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

func (p *easyPage) GetPeer() string {
	return p.conn.RemoteAddr().String()
}

func (p *easyPage) AddCss(css string) Page {
	id := util.GetCallerID(util.LevelParent)
	msg := toClientMsgData{id, "css", css}
	p.sendMsg(msg)
	return p
}

func (p *easyPage) Write(e any) string {
	id := util.GetCallerID(util.LevelParent)
	return p.WriteWithID(id, e)
}

func (p *easyPage) WriteWithID(id string, e any) string {
	msg := toClientMsgData{id, "", fmt.Sprint(e)}
	p.sendMsg(msg)
	if e, ok := e.(IEnableRegist); ok {
		e.RegistEvent(p)
	}
	return id
}

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
	toClient := toClientMsgData{id, "event", typ}
	p.sendMsg(toClient)
}

func (p *easyPage) Refresh(e IGetID) {
	id := e.GetID()
	p.WriteWithID(id, e)
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
					out := []byte{byte(CbDataTypeString)}
					out = append(out, []byte(msg.Msg)...)
					cb.Event.MessageCallbackFromFramwork(msg.ID, out)
				}
			case fileMsgData:
				cb := p.callback[msg.ID]
				if cb.Event != nil {
					out := []byte{byte(CbDataTypeBinary)}
					out = append(out, msg.BinaryData...)
					cb.Event.MessageCallbackFromFramwork(msg.ID, out)
				}
			case eventMsgData:
				p.callback[msg.ID] = msg
			default:
				fmt.Println("unknown msg type:", msg)
			}
		}
	}
}

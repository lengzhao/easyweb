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
	msg := ToClientMsgData{id, "title", title}
	p.sendMsg(msg)
	return p
}

func (p *easyPage) AddJs(js string) Page {
	id := util.GetCallerID(util.LevelParent)
	msg := ToClientMsgData{id, "js", js}
	p.sendMsg(msg)
	return p
}

func (p *easyPage) GetPeer() string {
	return p.conn.RemoteAddr().String()
}

func (p *easyPage) AddCss(css string) Page {
	id := util.GetCallerID(util.LevelParent)
	msg := ToClientMsgData{id, "css", css}
	p.sendMsg(msg)
	return p
}

func (p *easyPage) Write(e any) string {
	id := util.GetCallerID(util.LevelParent)
	return p.WriteWithID(id, e)
}

func (p *easyPage) WriteWithID(id string, e any) string {
	msg := ToClientMsgData{id, "", fmt.Sprint(e)}
	p.sendMsg(msg)
	if event, ok := e.(Event); ok {
		msg1 := ToClientMsgData{"", "event", ""}
		msg1.ID, msg1.Msg = event.EventInfo()
		p.sendMsg(msg1)
		msg2 := EventMsgData{}
		msg2.ID = msg1.ID
		msg2.E = event
		if fe, ok := e.(FileEvent); ok {
			fmt.Println("have filecb:", msg1.ID)
			msg2.F = fe
		}
		p.sendMsg(msg2)
	}
	return id
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
			case ToClientMsgData:
				if msg.ID == "" {
					continue
				}
				p.conn.WriteMessage(websocket.TextMessage, encode(msg))
				if msg.Msg == "" {
					delete(p.callback, msg.ID)
				}
			case FromClientMsgData:
				cb := p.callback[msg.ID]
				if cb.E != nil {
					cb.E.MessageCb(msg.ID, msg.Msg)
				}
			case FileMsgData:
				// fmt.Println("FileMsgData 01:", msg.ID, msg.File, msg.Size)
				cb := p.callback[msg.ID]
				if cb.F != nil {
					cb.F.FileCb(msg.ID, msg.File, msg.Size, msg.BinaryData)
				}
			case EventMsgData:
				p.callback[msg.ID] = msg
			default:
				fmt.Println("unknown msg type:", msg)
			}
		}
	}
}

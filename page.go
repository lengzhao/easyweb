package easyweb

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lengzhao/easyweb/util"
)

var _ Page = &easyPage{}

func (p *easyPage) Title(title string) Page {
	id := util.GetCallerID(util.LevelParent)
	msg := MsgData{id, "title", title, nil}
	select {
	case <-p.closed:
	case p.respChan <- msg:
	}
	return p
}

func (p *easyPage) AddJs(js string) Page {
	id := util.GetCallerID(util.LevelParent)
	msg := MsgData{id, "js", js, nil}
	select {
	case <-p.closed:
	case p.respChan <- msg:
	}
	return p
}

func (p *easyPage) GetPeer() string {
	return p.conn.RemoteAddr().String()
}

func (p *easyPage) AddCss(css string) Page {
	id := util.GetCallerID(util.LevelParent)
	msg := MsgData{id, "css", css, nil}
	select {
	case <-p.closed:
	case p.respChan <- msg:
	}
	return p
}

func (p *easyPage) Write(e any) string {
	id := util.GetCallerID(util.LevelParent)
	return p.WriteWithID(id, e)
}

func (p *easyPage) WriteWithID(id string, e any) string {
	msg := MsgData{id, "", fmt.Sprint(e), nil}
	msg1 := MsgData{"", "event", "", nil}
	if event, ok := e.(Event); ok {
		msg1.ID, msg1.Msg = event.EventInfo()
		msg1.cb = event
	}
	select {
	case <-p.closed:
		return ""
	case p.respChan <- msg:
		p.respChan <- msg1
	}
	return id
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

package easyweb

import (
	"github.com/gorilla/websocket"
)

type Page interface {
	Title(string) Page
	AddJs(string) Page
	AddCss(string) Page
	Write(any) string
	WriteWithID(string, any) string
	GetPeer() string
	Close()
	WaitUntilClosed()
}

type MsgData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Msg  string `json:"msg"`
	cb   Event
}

type easyPage struct {
	callback map[string]MessageCb
	conn     *websocket.Conn
	respChan chan MsgData
	reqChan  chan MsgData
	closed   chan int
}

type MessageCb interface {
	MessageCb(id, info string)
}
type PageFunc func(page Page)

type Event interface {
	EventInfo() (id, typ string)
	MessageCb(id, info string)
}

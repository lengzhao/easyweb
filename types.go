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

type ToClientMsgData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

type FromClientMsgData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

type FileMsgData struct {
	ID         string
	File       string
	Size       int64
	BinaryData []byte
}

type EventMsgData struct {
	ID string
	E  Event
	F  FileEvent
}

type easyPage struct {
	callback map[string]EventMsgData
	conn     *websocket.Conn
	closed   chan int
	msgChan  chan any
}

type MessageCb interface {
	MessageCb(id, info string)
}
type PageFunc func(page Page)

type Event interface {
	MessageCb
	EventInfo() (id, typ string)
}

type FileEvent interface {
	FileCb(id, fn string, size int64, data []byte)
}

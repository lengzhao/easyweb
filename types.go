package easyweb

import (
	"github.com/gorilla/websocket"
)

type Page interface {
	Title(string) Page
	AddJs(string) Page
	AddCss(string) Page
	RunJs(js string) Page
	Write(any) string
	WriteWithID(string, any) string
	Refresh(e IGetID)
	GetPeer() string
	Close()
	WaitUntilClosed()
	// regist element event after loaded
	RegistEvent(id, typ string, cb IMessageCb)
}

type toClientMsgData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

type fromClientMsgData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

type fileMsgData struct {
	ID         string
	File       string
	Size       int64
	BinaryData []byte
}

type eventMsgData struct {
	ID    string
	Event IMessageCb
}

type easyPage struct {
	callback map[string]eventMsgData
	conn     *websocket.Conn
	closed   chan int
	msgChan  chan any
}

type IGetID interface {
	GetID() string
}

type IMessageCb interface {
	MessageCallbackFromFramwork(id string, data []byte) bool
}

type PageFunc func(page Page)

type CbDataType byte

const (
	CbDataTypeString CbDataType = iota
	CbDataTypeBinary
)

type IAfterLoaded interface {
	AfterElementLoadedFromFramwork(p Page)
}

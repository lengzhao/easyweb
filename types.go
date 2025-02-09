package easyweb

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Page interface {
	Title(string) Page
	AddJs(string) Page
	AddCss(string) Page
	RunJs(js string) Page
	Write(any) string
	WriteWithID(string, any) string
	Replace(IGetID) string
	Delete(string)
	SetAttr(id, key, value string) string
	GetPeer() string
	Close()
	WaitUntilClosed()
	// regist element event after loaded
	RegistEvent(id, typ string, cb IMessageCb)

	SetEnv(key string, value any)
	GetEnv(key string) any
	WatchEnv(key string, cb func(value any)) error
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
	mu       sync.Mutex
	env      map[string]any
	watchEnv map[string]func(value any)
}

type IContainerID interface {
	ContainerID() string
}

type IMessageCb interface {
	MessageCallbackFromFramwork(page Page, id string, dataType CbDataType, data []byte) bool
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

const (
	ENV_COOKIES     = "e_cookies"     // http.Request.Cookies()
	ENV_PATH        = "e_path"        // http.Request.URL.Path
	ENV_HEADER      = "e_header"      // http.Request.Header
	ENV_QUERY       = "e_query"       // http.Request.URL.Query()
	ENV_CLOSEING    = "e_closing"     // bool,true
	ENV_REMOTE_ADDR = "e_remote_addr" // http.Request.RemoteAddr
)

type IGetID interface {
	GetID() string
}

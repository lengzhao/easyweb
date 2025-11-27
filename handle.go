package easyweb

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/lengzhao/easyweb/util"
)

//go:embed templates/default.html
var DefaultPageTmpl string

var WssPrefix = "/"
var DefaultPagePath []string = []string{
	"", "/", "/index.html",
}

// RegisterPage registers a page function to handle HTTP requests for the specified path patterns.
//
// Parameters:
// - pn: The page function to register.
// - path: Additional path patterns to register the page function for. if set to "", the caller's filename will be used. e.g. "index.go"-> "index.html"
//
// Return:
// - string: The WebSocket URL path for the registered page.
func RegisterPage(pn PageFunc, path ...string) string {
	pid := util.GetID()
	return RegisterPageWithID(pid, pn, path...)
}

func RegisterPageWithID(pid string, pn PageFunc, path ...string) string {
	pattern := make(map[string]bool)

	for _, it := range path {
		it = strings.TrimSpace(it)
		if it == "" {
			caller := util.GetCallerFile(util.LevelParent)
			caller = strings.Replace(caller, ".go", ".html", 1)
			it = "/" + caller
		}
		if strings.HasPrefix(it, "/") {
			pattern[it] = true
		} else {
			pattern["/"+it] = true
		}
	}

	page := newPageWs(pid, pn)
	for p := range pattern {
		page.addPath(p)
	}

	for p := range pattern {
		http.HandleFunc(p, page.HandleHtml)
		fmt.Println("register:", p, pid)
	}
	wssPath := "/wss/" + pid
	http.HandleFunc(wssPath, page.HandleWs)
	if WssPrefix != "" && wssPath != "/" {
		http.HandleFunc(WssPrefix+wssPath, page.HandleWs)
	}
	return wssPath
}

type pageWs struct {
	id       string
	pageData string
	cb       PageFunc
	paths    map[string]bool
	// Page实例，在服务启动时创建
	page Page
}

func newPageWs(id string, cb PageFunc) *pageWs {
	out := &pageWs{id, "", cb, make(map[string]bool), nil}
	// 在服务启动时创建Page实例
	out.page = newPage()
	buff := new(bytes.Buffer)
	t, err := template.New("1").Parse(DefaultPageTmpl)
	if err != nil {
		log.Println("fail to parse template:", err)
		return nil
	}
	t.Execute(buff, id)
	cb(out.page)
	out.pageData = buff.String()
	return out
}

func (p *pageWs) addPath(in ...string) {
	if p.paths == nil {
		p.paths = make(map[string]bool)
	}
	for _, it := range in {
		p.paths[it] = true
	}
}

// Session represents a user session
type session struct {
	callback map[string]eventMsgData
	page     *pageWs
	conn     *websocket.Conn
	closed   chan int
	msgChan  chan any
	mu       sync.Mutex
	env      map[string]any
	watchEnv map[string]func(value any)
}

func newSession(page *pageWs, conn *websocket.Conn) *session {
	var s session
	s.page = page
	s.conn = conn
	s.callback = make(map[string]eventMsgData)
	s.msgChan = make(chan any, 10)
	s.closed = make(chan int)
	s.env = make(map[string]any)
	s.watchEnv = make(map[string]func(value any))
	return &s
}

func (s *session) GetPeer() string {
	return s.conn.RemoteAddr().String()
}

func (s *session) Close() {
	select {
	case <-s.closed:
	default:
		s.conn.Close()
	}
}

func (s *session) WaitUntilClosed() {
	<-s.closed
}

func (s *session) sendMsg(msg any) {
	select {
	case <-s.closed:
	case s.msgChan <- msg:
	}
}

func (s *session) SetEnv(key string, value any) {
	s.mu.Lock()
	s.env[key] = value
	cb := s.watchEnv[key]
	s.mu.Unlock()
	if cb != nil {
		cb(value)
	}
}

func (s *session) GetEnv(key string) (value any) {
	s.mu.Lock()
	value = s.env[key]
	s.mu.Unlock()
	return
}

func (s *session) WatchEnv(key string, cb func(value any)) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.watchEnv[key] != nil {
		return fmt.Errorf("exist callback")
	}
	s.watchEnv[key] = cb
	return nil
}

func (s *session) Title(title string) Session {
	id := util.GetCallerID(util.LevelParent)
	msg := toClientMsgData{id, "title", title}
	s.sendMsg(msg)
	return s
}

func (s *session) AddJs(js string) Session {
	id := util.GetCallerID(util.LevelParent)
	msg := toClientMsgData{id, "js", js}
	s.sendMsg(msg)
	return s
}

func (s *session) RunJs(js string) Session {
	id := util.GetCallerID(util.LevelParent)
	msg := toClientMsgData{id, "eval", js}
	s.sendMsg(msg)
	return s
}

func (s *session) AddCss(css string) Session {
	id := util.GetCallerID(util.LevelParent)
	msg := toClientMsgData{id, "css", css}
	s.sendMsg(msg)
	return s
}

// e 将作为子元素放到div内，如<div id=rand_id>e</div>
func (s *session) Write(e any) string {
	var id string
	gid, ok := e.(IContainerID)
	if ok {
		id = gid.ContainerID()
	}
	msg, ok := e.(toClientMsgData)
	if ok {
		s.sendMsg(msg)
		return msg.ID
	}
	if id == "" {
		id = util.GetCallerID(util.LevelParent)
	}
	return s.WriteWithID(id, e)
}

// e 将作为子元素放到id所属的元素内，如<div id=id>e</div>
func (s *session) WriteWithID(id string, e any) string {
	if id == "" {
		id = util.GetCallerID(util.LevelParent)
	}
	log.Println("Session.WriteWithID", id, e)
	msg := toClientMsgData{id, "", fmt.Sprint(e)}
	s.sendMsg(msg)
	if e, ok := e.(IAfterLoaded); ok {
		e.AfterLoaded(s)
	}
	return id
}

func (s *session) Replace(e IGetID) string {
	id := e.GetID()
	if id == "" {
		return ""
	}
	msg := toClientMsgData{id, "replace", fmt.Sprint(e)}
	s.sendMsg(msg)
	return id
}

func (s *session) Delete(id string) {
	msg := toClientMsgData{id, "", ""}
	s.sendMsg(msg)
}

// 修改指定id的元素的属性
func (s *session) SetAttr(id, key, value string) string {
	info := attrInfo{Key: key, Value: value}
	data, _ := json.Marshal(info)
	msg := toClientMsgData{id, "attr", string(data)}
	s.sendMsg(msg)
	return id
}

// 注册指定id的前端事件，typ是前端事件名，cb是回调还是，如果cb为空，则是取消事件回调
func (s *session) RegistEvent(id, typ string, cb IMessageCb) {
	if id == "" || typ == "" {
		return
	}
	// 1. add event callback(server side)
	cbMsg := eventMsgData{}
	cbMsg.ID = id
	cbMsg.Event = cb
	s.sendMsg(cbMsg)
	// 2. add client event(jquery will handle the event)
	if cb != nil {
		toClient := toClientMsgData{id, "event", typ}
		s.sendMsg(toClient)
	} else {
		toClient := toClientMsgData{id, "off", typ}
		s.sendMsg(toClient)
	}
}

func (s *session) processMsg() {
	for {
		select {
		case <-s.closed:
			return
		case data, ok := <-s.msgChan:
			if !ok {
				continue
			}
			switch msg := data.(type) {
			case toClientMsgData:
				if msg.ID == "" {
					continue
				}
				s.conn.WriteMessage(websocket.TextMessage, encode(msg))
				if msg.Msg == "" {
					delete(s.callback, msg.ID)
				}
			case fromClientMsgData:
				if msg.Type == "ping" && msg.ID == "" {
					continue
				}
				log.Println("from client. id:", msg.ID, " msg:", msg.Msg, " type:", msg.Type)
				cb := s.callback[msg.ID]
				if cb.Event != nil {
					log.Println("call event callback of session", msg.ID, msg.Msg)
					go func(id string, dataType CbDataType, data []byte) {
						defer func() {
							if r := recover(); r != nil {
								fmt.Println("Recovered", r)
							}
						}()
						cb.Event.MessageCallbackFromFramwork(s, msg.ID, dataType, data)
					}(msg.ID, CbDataTypeString, []byte(msg.Msg))
				} else if s.page.page != nil {
					log.Println("call event callback of page", msg.ID, msg.Msg)
					// 如果Session没有匹配到事件处理器，交给Page处理
					go func(id string, dataType CbDataType, data []byte) {
						defer func() {
							if r := recover(); r != nil {
								fmt.Println("Recovered", r)
							}
						}()
						s.page.page.MessageCallbackFromFramwork(s, msg.ID, dataType, data)
					}(msg.ID, CbDataTypeString, []byte(msg.Msg))
				}
			case fileMsgData:
				cb := s.callback[msg.ID]
				if cb.Event != nil {
					go func(id string, dataType CbDataType, data []byte) {
						defer func() {
							if r := recover(); r != nil {
								fmt.Println("Recovered", r)
							}
						}()
						cb.Event.MessageCallbackFromFramwork(s, msg.ID, dataType, data)
					}(msg.ID, CbDataTypeBinary, msg.BinaryData)
				} else if s.page.page != nil {
					// 如果Session没有匹配到事件处理器，交给Page处理
					go func(id string, dataType CbDataType, data []byte) {
						defer func() {
							if r := recover(); r != nil {
								fmt.Println("Recovered", r)
							}
						}()
						s.page.page.MessageCallbackFromFramwork(s, msg.ID, dataType, data)
					}(msg.ID, CbDataTypeBinary, msg.BinaryData)
				}
			case eventMsgData:
				if msg.Event == nil {
					delete(s.callback, msg.ID)
				} else {
					s.callback[msg.ID] = msg
				}
			default:
				fmt.Println("unknown msg type:", msg)
			}
		}
	}
}

var upgrader = websocket.Upgrader{} // use default options
func (p *pageWs) HandleWs(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("ws connect:", r.RemoteAddr)
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	// Page在服务启动时已创建，这里只创建Session
	session := newSession(p, c)
	go session.processMsg()

	session.SetEnv(ENV_REMOTE_ADDR, r.RemoteAddr)
	session.SetEnv(ENV_HEADER, r.Header)
	session.SetEnv(ENV_QUERY, r.URL.Query())
	session.SetEnv(ENV_PATH, r.URL.Path)
	session.SetEnv(ENV_COOKIES, r.Cookies())

	p.page.PageLoaded(session)
	var lastID string
	var lastFile string
	var lastSize int64
	for {
		msgType, data, err := c.ReadMessage()
		// fmt.Println("ReadMessage:", msgType, len(data), err)
		if err != nil {
			break
		}
		if msgType == websocket.TextMessage {
			var msg fromClientMsgData
			err = json.Unmarshal(data, &msg)
			if err != nil {
				continue
			}
			// 将事件信息传递给Session处理
			session.msgChan <- msg
			continue
		}
		if msgType != websocket.BinaryMessage {
			continue
		}
		if lastFile == "" && bytes.HasPrefix(data[:20], []byte("file:")) {
			arr := strings.Split(string(data), ":")
			if len(arr) != 5 {
				continue
			}
			fmt.Println("data:", len(data), string(data), len(string(data)))
			lastID = arr[1]
			lastFile = arr[2]
			lastSize, _ = strconv.ParseInt(arr[3], 10, 64)
			continue
		}
		var msg fileMsgData
		msg.ID = lastID
		msg.File = lastFile
		msg.Size = lastSize
		msg.BinaryData = data
		lastFile = ""
		lastID = ""
		// 将事件信息传递给Session处理
		session.msgChan <- msg
		if len(data) != int(lastSize) {
			fmt.Println("fileMsgData,wrong size:", msg.ID, msg.File, "hope size:", msg.Size, "data size:", len(data))
		}

	}
	session.SetEnv(ENV_CLOSEING, true)
	close(session.closed)
	// fmt.Println("ws disconnect:", r.RemoteAddr)
}

func (p *pageWs) HandleHtml(w http.ResponseWriter, r *http.Request) {
	fn := path.Clean(r.URL.Path)
	if !p.paths[fn] {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, p.pageData)
}

func encode(v interface{}) []byte {
	buff := new(strings.Builder)
	enc := json.NewEncoder(buff)
	enc.SetEscapeHTML(false)
	enc.Encode(v)
	return []byte(buff.String())
}

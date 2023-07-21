package easyweb

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/lengzhao/easyweb/util"
)

type Page interface {
	Title(string) Page
	AddJs(string) Page
	AddCss(string) Page
	Write(any) string
}

type MsgData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

type chanData struct {
	ID   string
	Type string
	Msg  any
}

type easyPage struct {
	elements map[string]any
	conn     *websocket.Conn
	respChan chan chanData
	reqChan  chan MsgData
	closed   chan int
}

type MsgType string

const (
	MsgTypeNull   MsgType = ""
	MsgTypeInput  MsgType = "input"
	MsgTypeButton MsgType = "button"
	MsgTypeForm   MsgType = "form"
)

type GetID interface {
	GetID() string
}
type GetType interface {
	GetType() string
}
type ElementCb interface {
	ElementCb(id, info string)
}
type PageFunc func(page Page)

type pageInfo struct {
	title string
	cb    PageFunc
}

//go:embed templates
var pageTemplate embed.FS

const IndexPage = "/index.html"

func foundFromEmbed(w http.ResponseWriter, r *http.Request, fn string) error {
	// fmt.Println("fn:", fn)
	if fn == "." || fn == "/" {
		fn = IndexPage
	}
	fn = "templates" + fn
	// fmt.Println("fn2:", fn)
	f, err := pageTemplate.Open(fn)
	if err != nil {
		// fmt.Println("not found file:", fn)
		return err
	}
	defer f.Close()
	io.Copy(w, f)
	return nil
}

func init() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fn := path.Clean(r.URL.Path)
		err := foundFromEmbed(w, r, fn)
		if err == nil {
			return
		}
		http.NotFound(w, r)
	}))
}

func NewPage(pagePath string, cb PageFunc) {
	fn := pagePath
	if fn == "" {
		fn = util.GetCallerFile(util.LevelParent)
		fn = strings.Replace(fn, ".go", ".html", 1)
		// fmt.Println("fn:", fn)
	}
	if !strings.HasPrefix(fn, "/") {
		fn = "/" + fn
	}
	if fn == "/" {
		fn = IndexPage
	}

	var p pageInfo
	p.title = fn
	p.cb = cb
	http.HandleFunc(fn, func(w http.ResponseWriter, r *http.Request) {
		foundFromEmbed(w, r, IndexPage)
	})
	http.HandleFunc("/ws"+fn, p.echo)
	if fn == IndexPage {
		http.HandleFunc("/ws", p.echo)
	}
}

var upgrader = websocket.Upgrader{} // use default options
var _ Page = &easyPage{}

func (p pageInfo) echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	var page easyPage
	page.conn = c
	page.elements = make(map[string]any)
	page.respChan = make(chan chanData)
	page.reqChan = make(chan MsgData, 10)
	page.closed = make(chan int)
	go func() {
		// defer func() { recover() }()
		p.cb(&page)
	}()
	go func() {
		for {
			select {
			case <-page.closed:
				return
			case msg, ok := <-page.respChan:
				if !ok {
					return
				}
				page.elements[msg.ID] = msg.Msg
				respData := MsgData{ID: msg.ID, Type: msg.Type, Msg: fmt.Sprint(msg.Msg)}
				page.conn.WriteMessage(websocket.TextMessage, encode(respData))
			case msg, ok := <-page.reqChan:
				if !ok {
					continue
				}
				e := page.elements[msg.ID]
				cb, ok := e.(ElementCb)
				// fmt.Println("cb type:", msg.ID, cb, ok, msg.Type)
				if ok {
					cb.ElementCb(msg.ID, msg.Msg)
				}
			}
		}
	}()
	for {
		_, data, err := c.ReadMessage()
		if err != nil {
			break
		}
		var msg MsgData
		err = json.Unmarshal(data, &msg)
		if err != nil {
			continue
		}

		log.Println("id:", msg.ID, msg.Type, msg.Msg)
		page.reqChan <- msg
	}
	close(page.closed)
}

func (p *easyPage) Title(title string) Page {
	id := util.GetCallerID(util.LevelParent)
	msg := chanData{id, "title", title}
	select {
	case <-p.closed:
	case p.respChan <- msg:
	}
	return p
}

func (p *easyPage) AddJs(js string) Page {
	id := util.GetCallerID(util.LevelParent)
	msg := chanData{id, "js", js}
	select {
	case <-p.closed:
	case p.respChan <- msg:
	}
	return p
}

func (p *easyPage) AddCss(css string) Page {
	id := util.GetCallerID(util.LevelParent)
	msg := chanData{id, "css", css}
	select {
	case <-p.closed:
	case p.respChan <- msg:
	}
	return p
}

func (p *easyPage) Write(e any) string {
	id := ""
	if getId, ok := e.(GetID); ok {
		id = getId.GetID()
	}
	if id == "" {
		id = util.GetCallerID(util.LevelParent)
	}
	msg := chanData{id, "", e}
	if getType, ok := e.(GetType); ok {
		msg.Type = getType.GetType()
	}
	select {
	case <-p.closed:
		return ""
	case p.respChan <- msg:
	}
	return id
}

func encode(v interface{}) []byte {
	buff := new(strings.Builder)
	enc := json.NewEncoder(buff)
	enc.SetEscapeHTML(false)
	enc.Encode(v)
	return []byte(buff.String())
}

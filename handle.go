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
	"strings"

	"github.com/gorilla/websocket"
	"github.com/lengzhao/easyweb/util"
)

//go:embed templates/default.html
var DefaultPageTmpl string

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
	pattern := make(map[string]bool)
	pid := util.GetID()

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
	return wssPath
}

type pageWs struct {
	id       string
	pageData string
	cb       PageFunc
	paths    map[string]bool
}

func newPageWs(id string, cb PageFunc) *pageWs {
	out := &pageWs{id, "", cb, make(map[string]bool)}
	buff := new(bytes.Buffer)
	t, err := template.New("1").Parse(DefaultPageTmpl)
	if err != nil {
		log.Println("fail to parse template:", err)
		return nil
	}
	t.Execute(buff, id)
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

var upgrader = websocket.Upgrader{} // use default options
func (p *pageWs) HandleWs(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("ws connect:", r.RemoteAddr)
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	var page easyPage
	page.conn = c
	page.callback = make(map[string]MessageCb)
	page.respChan = make(chan MsgData)
	page.reqChan = make(chan MsgData, 10)
	page.closed = make(chan int)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered", r)
			}
		}()
		p.cb(&page)
	}()
	go func() {
		for {
			select {
			case <-page.closed:
				return
			case msg, ok := <-page.respChan:
				if !ok || msg.ID == "" {
					continue
				}
				if msg.Type == "event" {
					page.callback[msg.ID] = msg.cb
				}
				page.conn.WriteMessage(websocket.TextMessage, encode(msg))
			case msg, ok := <-page.reqChan:
				if !ok {
					continue
				}
				cb := page.callback[msg.ID]
				if cb != nil {
					cb.MessageCb(msg.ID, msg.Msg)
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
		// log.Println("id:", msg.ID, msg.Type, msg.Msg)
		page.reqChan <- msg
	}
	close(page.closed)
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

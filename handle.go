package easyweb

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
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
	http.HandleFunc("/wss/"+pid+"/upload/", page.HandleUpload)
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
	page.callback = make(map[string]eventMsgData)
	page.msgChan = make(chan any, 10)
	page.closed = make(chan int)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered", r)
			}
		}()
		p.cb(&page)
	}()
	go page.processMsg()
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
			page.msgChan <- msg
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
		page.msgChan <- msg
		if len(data) != int(lastSize) {
			fmt.Println("fileMsgData,wrong size:", msg.ID, msg.File, "hope size:", msg.Size, "data size:", len(data))
		}

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

func (p *pageWs) HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	fn := path.Clean(r.URL.Path)
	r.ParseMultipartForm(32 << 20)

	msgInfo := make(map[string]string)
	for k, v := range r.Form {
		msgInfo[k] = strings.Join(v, ",")
	}

	for k := range r.MultipartForm.File {
		_, fileHeader, err := r.FormFile(k)
		if err != nil {
			fmt.Println(k, err)
			continue
		}
		msgInfo[k] = fileHeader.Filename
	}
	fData, _ := json.MarshalIndent(msgInfo, "", "  ")
	fmt.Println("form data:", fn, string(fData))

	file, handler, err := r.FormFile("file1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

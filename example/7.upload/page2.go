package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/e"
)

func page2() {
	var uploadPath string = "/upload/123"
	http.HandleFunc(uploadPath, upload)

	easyweb.RegisterPage(func(page easyweb.Page) {
		page.Title("MyWeb")

		page.Write("<h3>Will upload file and save it to ./test/ on server</h3>")
		page.Write(e.Link("home", "/"))

		fileInput := e.InputGroup("file1", "File").ChangeType("file")

		form := e.Form(func(id string, info map[string]string) {
			log.Panicln("hope not any event")
		})
		form.Attr("id", "")
		form.Add(fileInput).Action(uploadPath, "")
		form.SetFileCb(func(id string, data []byte) {
			log.Panicln("hope not any file")
		})
		page.Write(form)
	}, "/page2.html")
}

// upload logic
func upload(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("method:", r.Method)
	if r.Method != http.MethodPost {
		return
	}
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("file1")
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "fail to get file1:", err)
		return
	}
	defer file.Close()
	// fmt.Fprintf(w, "%v.", handler.Header)
	os.Mkdir("./test", 0766)
	f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "fail to OpenFile:", err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprint(w, "success upload, save file to ./test/"+handler.Filename+" on server")
}

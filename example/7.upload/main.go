package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/e"
)

// upload logic
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method != http.MethodPost {
		return
	}
	r.ParseMultipartForm(32 << 20)
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

func main() {
	http.HandleFunc("/upload", upload)

	easyweb.RegisterPage(func(page easyweb.Page) {
		page.Title("MyWeb")
		radio := e.Radio("radio1").Class("col-3").Add(map[string]string{"r1": "r1", "r2": "r2", "r3": "r3"})
		selectItem := e.Select("select1").Add(e.SelectItem{Value: "s1", Text: "s1", Selected: false})
		selectItem.Add(e.SelectItem{Value: "s2", Text: "s2", Selected: true})
		selectItem.Add(e.SelectItem{Value: "s3", Text: "s3", Selected: false})
		fileInput := e.Entry("file1").SetType("file").Prefix("File")

		form := e.Form(func(id string, info map[string]string) {
			fmt.Println("form data:", info)
		})
		form.Add(radio).Add(selectItem).Add(fileInput)
		form.SetFileCb(func(id, fn string, size int64, data []byte) {
			fmt.Println("FileCb", id, fn, size, len(data))
		})
		page.Write(form)
	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)

}

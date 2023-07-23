package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/e"
)

func main() {
	page2()

	easyweb.RegisterPage(func(page easyweb.Page) {
		page.Title("MyWeb")

		fileInput := e.Entry("file1").SetType("file").Prefix("File")

		page.Write(e.Link("page2.html: upload to other path", "/page2.html"))
		var statID string
		form := e.Form(func(id string, info map[string]string) {
			fmt.Println("form data:", info)
			page.Write(e.Label("receive form data:").Write(info))
		})
		form.Add(fileInput)
		form.SetFileCb(func(id, fn string, size int64, data []byte) {
			fmt.Println("FileCb", id, fn, size, len(data))
			page.WriteWithID(statID, e.Label("stat:finish. file:"+fn+" size:"+strconv.FormatInt(size, 10)+" bytes"))
		})
		page.Write(form)
		statID = page.Write(e.Label("stat:"))
	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)

}

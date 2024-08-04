package main

import (
	"fmt"
	"net/http"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/e"
)

func main() {
	page2()

	easyweb.RegisterPage(func(page easyweb.Page) {
		page.Title("MyWeb")

		fileInput := e.InputGroup("file1", "File").ChangeType("file")

		page.Write(e.Link("page2.html: upload to other path", "/page2.html"))
		var statID string
		form := e.Form(func(p easyweb.Page, id string, info map[string]string) {
			fmt.Println("form data:", info)
			page.Write(e.Label("receive form data:").Add(info))
		})
		form.AddAny(fileInput)
		form.SetFileCb(func(p easyweb.Page, id string, data []byte) {
			fmt.Println("FileCb", id, len(data))
			page.WriteWithID(statID, e.Label(fmt.Sprintf("receive file. id: %s. size:%d", id, len(data))))
		})
		page.Write(form)
		statID = page.Write(e.Label("stat:"))
	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)

}

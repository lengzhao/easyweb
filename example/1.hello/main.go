package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/e"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	easyweb.RegisterPage(func(page easyweb.Page) {
		page.Title("MyWeb")
		page.Write("<h1>Hello World</h1>")
		page.AddCss("/static/css/format2.css")
		page.Write(e.InputGroup("333", "@BBB").Suffix("$$$$"))
		var count int
		page.Write(e.Button("Click", func(id string) {
			fmt.Println("button click001:", id)
			count++
			page.Write("button click:" + fmt.Sprint(count))
			page.Write(time.Now().String())
		}))
	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)

}

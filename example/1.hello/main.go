package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/e"
)

func main() {
	easyweb.NewPage(easyweb.IndexPage, func(page easyweb.Page) {
		page.Title("MyWeb")
		page.Write(e.Navbar("MyWeb").Write(map[string]string{
			"Home":     "#",
			"Features": "#",
			"Pricing":  "#",
		}))
		page.Write("this is my first ui.")
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
			e := e.Label("-----").Write(i).Write("***<div>kkk</div>***")
			page.Write(e)
		}
		page.Write("fdsfdsfdsfdsfdsfsdfdsfdsfd")
		page.AddCss("/static/css/format2.css")
		page.Write(e.Entry("222").Prefix("@AAAA"))
		page.Write(e.Entry("333").Prefix("@BBB").Suffix("$"))
		page.Write(e.Button("Click").SetCb(func(id string) {
			fmt.Println("button click001:", id)
		}))
	})
	http.ListenAndServe(":8182", nil)

}

package main

import (
	"net/http"
	"time"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/e"
)

func main() {
	easyweb.NewPage(easyweb.IndexPage, func(page easyweb.Page) {
		page.Title("MyWeb")
		page.Write("<h1>H1</h1>")
		page.Write("this is my first ui.")
		page.Write("<a href=\"main.html\">Second Page</a>")
		for i := 0; i < 20; i++ {
			time.Sleep(100 * time.Millisecond)
			page.Write(e.Label("----").Write(i))
		}
		page.Write("fdsfdsfdsfdsfdsfsdfdsfdsfd")
	})
	easyweb.NewPage("", func(page easyweb.Page) {
		page.Title("MyWeb")
		page.Write(e.Label("<h1>Page2</h1>"))
		page.Write("this is my second page.")
		page.Write("<a href=\"index.html\">Home Page</a>")
	})
	http.ListenAndServe(":8182", nil)

}

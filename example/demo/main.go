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
		page.Title("Demo")
		tabs := e.Tabs()
		tabs.Add("Label", e.Box("<h1>Hello World</h1>", "<h2>Hello World</h2>", "<h3>Hello World</h3>"))
		tabs.Add("Accordion", e.Accordion().AddItem("H1", "<h1>Hello World</h1>").AddItem("H2", "<h2>Hello World</h2>").AddItem("H3", "<h3>Hello World</h3>").AddItem("H3-2", "<h3>Hello World</h3>"))
		tabs.Add("SVG", `<svg xmlns="http://www.w3.org/2000/svg" version="1.1">
							<circle cx="100" cy="50" r="40" stroke="black"
							stroke-width="2" fill="red" />
						</svg>`)

		tabs.Add("Entry", e.Entry("333").Prefix("@BBB").Suffix("$$$$").Value("hello"))
		var count int
		tabs.Add("Button", e.Button("Click", func(id string) {
			fmt.Println("button click001:", id)
			page.Write("button click:" + fmt.Sprint(count))
			page.Write(time.Now().String())
		}))

		page.Write(tabs)

	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)

}

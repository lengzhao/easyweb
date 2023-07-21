package main

import (
	"net/http"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/e"
)

func main() {
	easyweb.NewPage(easyweb.IndexPage, func(page easyweb.Page) {
		page.Title("MyWeb")
		page.Write(e.Entry("01").Value("hello"))
		radio := e.Radio("radio1").Class("ml-3").Add(map[string]string{"r1": "r1", "r2": "r2", "r3": "r3"})
		selectItem := e.Select("select1").Add(e.SelectItem{Value: "s1", Text: "s1", Selected: false})
		selectItem.Add(e.SelectItem{Value: "s2", Text: "s2", Selected: true})
		selectItem.Add(e.SelectItem{Value: "s3", Text: "s3", Selected: false})
		page.Write(e.Form().Add(radio).Add(selectItem))
	})
	http.ListenAndServe(":8182", nil)

}

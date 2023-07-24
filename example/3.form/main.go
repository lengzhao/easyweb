package main

import (
	"fmt"
	"net/http"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/e"
)

func main() {
	easyweb.RegisterPage(func(page easyweb.Page) {
		page.Title("MyWeb")
		// page.Write(e.Entry("01").Value("hello"))
		radio := e.Radio("radio1").Add(map[string]string{"name1": "text1", "name2": "text2", "name3": "text3"}).Inline()
		selectItem := e.Select("select1").Add(e.SelectItem{Value: "s1", Text: "s1", Selected: false})
		selectItem.Add(e.SelectItem{Value: "s2", Text: "s2", Selected: true})
		selectItem.Add(e.SelectItem{Value: "s3", Text: "s3", Selected: false})
		page.Write(e.Form(func(id string, info map[string]string) {
			fmt.Println("form data:", info)
		}).Add(radio).Add(selectItem))
	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)

}

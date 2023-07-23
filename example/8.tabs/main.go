package main

import (
	"fmt"
	"net/http"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/e"
)

func main() {
	easyweb.RegisterPage(func(page easyweb.Page) {
		page.Title("Tabs")
		tabs := e.Tabs()
		radio := e.Radio("radio1").Class("col-3").Add(map[string]string{"r1": "r1", "r2": "r2", "r3": "r3"})
		selectItem := e.Select("select1").Add(e.SelectItem{Value: "s1", Text: "s1", Selected: false})
		selectItem.Add(e.SelectItem{Value: "s2", Text: "s2", Selected: true})
		selectItem.Add(e.SelectItem{Value: "s3", Text: "s3", Selected: false})
		stat := e.Box("stat:")
		form := e.Form(func(id string, info map[string]string) {
			fmt.Println("form data:", info)
			stat.Add("form data:", info)
			page.WriteWithID(stat.GetID(), stat)
		}).Add(radio).Add(selectItem)
		tabs.Add("tab1", e.Box(form, stat))
		tabs.Add("tab2", e.Button("button1", nil))
		tabs.Add("tab3", "<h1>h1</h1>")
		page.Write(tabs)
		page.WriteWithID(form.GetID(), form)
	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)

}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/e"
)

func main() {
	easyweb.RegisterPage(func(page easyweb.Page) {
		page.Title("Tabs")
		tabs := e.Tabs()
		radio := e.Radio("radio1").AddItem("name1", "text1").AddItem("name2", "text2").AddItem("name3", "text3").Inline().Select("name2")
		selectItem := e.Select("select1").AddItem("s1", "s1").AddItem("s2", "s2").AddItem("s3", "s3").Select("s2")
		stat := e.Box(e.Div("stat:"))
		form := e.Form(func(p easyweb.Page, id string, info map[string]string) {
			fmt.Println("form data:", info)
			data, _ := json.MarshalIndent(info, "", "  ")
			stat.Add(e.Div("form data:"), e.Div(string(data)))
			page.WriteWithID(stat.GetID(), stat)
		}).AddItem(radio).AddItem(selectItem)
		tabs.AddItem("tab1", e.Box(form, stat))
		tabs.AddItem("tab2", e.Button("button1", nil))
		tabs.AddItem("tab3", e.MustParseHtml("<h1>h1</h1>"))
		page.Write(tabs)
		page.WriteWithID(form.GetID(), form)
	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)

}

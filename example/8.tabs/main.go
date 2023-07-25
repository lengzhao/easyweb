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
		radio := e.Radio("radio1").Add("name1", "text1").Add("name2", "text2").Add("name3", "text3").Inline().Check("name2")
		selectItem := e.Select("select1").Add("s1", "s1").Add("s2", "s2").Add("s3", "s3").Select("s2")
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

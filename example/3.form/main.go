package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/e"
)

func main() {
	easyweb.RegisterPage(func(page easyweb.Page) {
		page.Title("MyWeb")
		// page.Write(e.Entry("01").Value("hello"))
		radio := e.Radio("radio1").AddItem("name1", "text1").AddItem("name2", "text2").AddItem("name3", "text3").Inline().Select("name2")
		selectItem := e.Select("select1").AddItem("s1", "s1").AddItem("s2", "s2").AddItem("s3", "s3").Select("s2")
		page.Write(e.Form(func(p easyweb.Page, id string, info map[string]string) {
			fmt.Println("form data:", info)
			page.Write("form data:")
			page.Write(info)
			page.Write(time.Now().String())
		}).AddItem(radio).AddItem(selectItem).AddInput("name8", "Name8").AddItem(e.Textarea("textarea", "Textarea")))
	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)

}

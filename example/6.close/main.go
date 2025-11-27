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
		page.Write(e.InputGroup("01", "Hello").Value("hello"))
		radio := e.Radio("radio1").AddItem("name1", "text1").AddItem("name2", "text2").AddItem("name3", "text3").Inline()
		selectItem := e.Select("select1").AddItem("s1", "v1").AddItem("s2", "v2").AddItem("s3", "v3").Select("s2")
		page.Write(e.Form(func(p easyweb.Session, id string, info map[string]string) {
			fmt.Println("form data:", info)
		}).AddItem(radio).AddItem(selectItem))
		for i := 0; i < 5; i++ {
			page.Write(fmt.Sprintf("the websocket will be closed in %d seconds", 5-i))
			time.Sleep(1 * time.Second)
		}
		page.Write("the websocket is closed")
		fmt.Println("finish")
	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)

}

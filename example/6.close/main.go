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
		page.Write(e.Entry("01").Value("hello"))
		radio := e.Radio("radio1").Class("col-3").Add(map[string]string{"r1": "r1", "r2": "r2", "r3": "r3"})
		selectItem := e.Select("select1").Add(e.SelectItem{Value: "s1", Text: "s1", Selected: false})
		selectItem.Add(e.SelectItem{Value: "s2", Text: "s2", Selected: true})
		selectItem.Add(e.SelectItem{Value: "s3", Text: "s3", Selected: false})
		page.Write(e.Form(func(id string, info map[string]string) {
			fmt.Println("form data:", info)
		}).Add(radio).Add(selectItem))
		page.Write("the websocket will be close in 2 seconds")
		time.Sleep(2 * time.Second)
		page.Write("the websocket is closed")
		page.Close()
		page.WaitUntilClosed()
		fmt.Println("finish")
	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)

}

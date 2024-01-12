package main

import (
	"net/http"
	"time"

	"github.com/lengzhao/easyweb"
)

func main() {
	easyweb.RegisterPage(func(page easyweb.Page) {
		page.Title("MyWeb")
		page.Write("<h1>easy clock</h1>")
		page.WriteWithID("did", "this will be deleted after 10 second")
		t := time.Tick(time.Second)
		var id string
		for i := 0; i < 10; i++ {
			id = page.Write(time.Now().Local().String())
			<-t
		}
		page.WriteWithID(id, "the clock is stoped")
		page.Delete("did")
	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)

}

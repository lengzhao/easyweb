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
		t := time.Tick(time.Second)
		for {
			<-t
			id := page.Write(time.Now().Local().String())
			if id == "" {
				break
			}
		}
	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)

}

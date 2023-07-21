package main

import (
	"net/http"
	"time"

	"github.com/lengzhao/easyweb"
)

func main() {
	easyweb.NewPage(easyweb.IndexPage, func(page easyweb.Page) {
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
	})
	http.ListenAndServe(":8182", nil)

}

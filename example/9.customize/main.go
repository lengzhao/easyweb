package main

import (
	"net/http"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/e"
)

func main() {
	easyweb.RegisterPage(func(page easyweb.Page) {
		page.Title("MyWeb")
		page.Write("<h1>Hello World</h1>")

		item, _ := e.ParseHtml(`<div class="dropdown">
			<button class="btn btn-secondary dropdown-toggle" type="button" id="dropdownMenuButton1" data-bs-toggle="dropdown" aria-expanded="false">
			  Dropdown button
			</button>
			<ul class="dropdown-menu" aria-labelledby="dropdownMenuButton1">
			  <li><a class="dropdown-item" href="#">Action</a></li>
			  <li><a class="dropdown-item" href="#">Another action</a></li>
			  <li><a class="dropdown-item" href="#">Something else here</a></li>
			</ul>
		  </div>`)
		page.Write(item)

	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)

}

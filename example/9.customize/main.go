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

		item2, _ := e.ParseHtml(`<nav class="navbar navbar-expand-md navbar-light bg-light">
		<a class="navbar-brand" href="#">Navbar</a>
		<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
		  <span class="navbar-toggler-icon"></span>
		</button>
		<div class="collapse navbar-collapse" id="navbarNav">
		  <ul class="navbar-nav">
			<li class="nav-item active">
			  <a class="nav-link" href="#">Home <span class="sr-only">(current)</span></a>
			</li>
			<li class="nav-item">
			  <a class="nav-link" href="#">Features</a>
			</li>
			<li class="nav-item">
			  <a class="nav-link" href="#">Pricing</a>
			</li>
			<li class="nav-item">
			  <a class="nav-link disabled" href="#">Disabled</a>
			</li>
		  </ul>
		</div>
	  </nav>
	  `)
		page.Write(item2)

	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)

}

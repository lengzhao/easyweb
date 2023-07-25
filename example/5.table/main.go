package main

import (
	"net/http"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/e"
)

func main() {
	easyweb.RegisterPage(func(page easyweb.Page) {
		page.Title("MyWeb")
		page.Write("<h1>Table</h1>")
		table := e.Table("First", "Last", "Handle").BoldFirstRow()
		table.AddValue([][]any{
			{"Mark", "Otto", "@mdo"},
			{"Jacob", "Thornton", "@fat"},
			{"Larry", "the Bird", "@twitter"},
			{"Larry", "the Bird"},
		})
		page.Write(table)
	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)
}

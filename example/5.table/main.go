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
		page.Write(e.Button("UpdateTable", func(id string) {
			table.AddValue([][]any{
				{"Mark", "Otto", "@mdo"},
			})
			table.Refresh(page)
		}))

		page.Write("<h1>Map Table</h1>")
		page.Write(e.Map2Table(3, map[string]any{
			"First":  "Mark",
			"Last":   "Otto",
			"Handle": "@mdo",
			"Others": []any{"Jacob", "Thornton", "@fat", "Larry", "the Bird", "@twitter"},
			"Map": map[string]any{
				"First":  "Mark",
				"Last":   "Otto",
				"Handle": "@mdo",
			},
		}))
	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)
}

package e

import (
	_ "embed"
)

type elementSidebar struct {
	HtmlToken
}

//go:embed templates/sidebar.html
var sidebar2 string

func Sidebar() *elementSidebar {
	out := elementSidebar{}
	out.Parse(sidebar2)
	out.Attr("id", getID())
	return &out
}

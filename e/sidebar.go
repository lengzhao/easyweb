package e

import (
	_ "embed"
)

type elementSidebar struct {
	HtmlToken
	id string
}

var _ IElement = &elementSidebar{}

func Sidebar(title string) *elementSidebar {
	id := "offcanvas" + getID()
	out := elementSidebar{}
	out.id = id
	out.parseText(`<div class="offcanvas offcanvas-start" data-bs-scroll="true" data-bs-backdrop="false" tabindex="-1" id="` + id + `" aria-labelledby="` + id + `Label">
	<div class="offcanvas-header">
	  <h5 class="offcanvas-title" id="` + id + `Label">` + title + `</h5>
	  <button type="button" class="btn-close text-reset" data-bs-dismiss="offcanvas" aria-label="Close"></button>
	</div>
	<div class="offcanvas-body">
	</div>
  </div>`)
	return &out
}

func (e *elementSidebar) GetButton() IElement {
	item, _ := ParseHtml(`<button class="btn btn-primary" type="button" data-bs-toggle="offcanvas" data-bs-target="#` + e.id + `" aria-controls="` + e.id + `">Enable body scrolling</button>`)
	return item
}

func (e *elementSidebar) AddItem(in IElement) *elementSidebar {
	e.children[1].Add(in)
	return e
}

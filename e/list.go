package e

import "fmt"

type listElement struct {
	HtmlToken
}

var _ IElement = &listElement{}

func List(in ...any) *listElement {
	var out listElement
	out.parseText(`<ul class="list-group"></ul>`)
	out.SetAttr("id", getID())
	out.AddItems(in...)
	return &out
}

func (e *listElement) ShowIndex() *listElement {
	e.SetAttr("class", e.GetAttr("class")+" list-group-numbered")
	return e
}

func (e *listElement) Horizontal() *listElement {
	e.SetAttr("class", e.GetAttr("class")+" list-group-horizontal")
	return e
}

func (b *listElement) AddItems(in ...any) *listElement {
	for _, v := range in {
		item := MustParseHtml(`<li class="list-group-item"></li>`)
		switch it := v.(type) {
		case IElement:
			item.Add(it)
		case string:
			item = MustParseHtml(`<li class="list-group-item">` + it + `</li>`)
		default:
			str := fmt.Sprint(it)
			item = MustParseHtml(`<li class="list-group-item">` + str + `</li>`)
		}
		b.Add(item)
	}
	return b
}

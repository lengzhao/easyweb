package e

type listElement struct {
	HtmlToken
}

func List(in ...any) *listElement {
	var out listElement
	out.parseText(`<ul class="list-group"></ul>`)
	out.Add(in...)
	return &out
}

func (e *listElement) ShowIndex() *listElement {
	e.Attr("class", e.GetAttr("class")+" list-group-numbered")
	return e
}

func (e *listElement) Horizontal() *listElement {
	e.Attr("class", e.GetAttr("class")+" list-group-horizontal")
	return e
}

func (b *listElement) Add(in ...any) *listElement {
	for _, v := range in {
		item, _ := ParseHtml(`<li class="list-group-item"></li>`)
		item.add(v)
		b.add(item)
	}
	return b
}

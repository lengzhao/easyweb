package e

type listElement struct {
	HtmlToken
}

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
		item, _ := ParseHtml(`<li class="list-group-item"></li>`)
		item.AddAny(v)
		b.Add(item)
	}
	return b
}

package e

type linkElement struct {
	HtmlToken
}

func Link(text, url string) *linkElement {
	var out linkElement
	out.parseText(`<a href="` + url + `">` + text + `</a>`)
	out.SetAttr("id", getID())
	return &out
}

func (b *linkElement) Blank() *linkElement {
	b.SetAttr("target", "_blank")
	return b
}

func (b *linkElement) AsButton() *linkElement {
	b.SetAttr("class", "btn btn-primary")
	b.SetAttr("role", "button")
	return b
}

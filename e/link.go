package e

type linkElement struct {
	HtmlToken
}

func Link(text, url string) *linkElement {
	var out linkElement
	out.parseText(`<a href="` + url + `">` + text + `</a>`)
	return &out
}

func (b *linkElement) Blank() *linkElement {
	b.Attr("target", "_blank")
	return b
}

func (b *linkElement) AsButton() *linkElement {
	b.Attr("class", "btn btn-primary")
	b.Attr("role", "button")
	return b
}

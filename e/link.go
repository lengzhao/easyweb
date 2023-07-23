package e

type LinkElement struct {
	// BaseElement
	node *HtmlNode
}

func Link(text, url string) *LinkElement {
	var out LinkElement
	out.node = NewNode("a")
	out.node.SetText(text)
	out.node.SetAttr("href", url)
	return &out
}

func (b *LinkElement) Blank() *LinkElement {
	b.node.SetAttr("target", "_blank")
	return b
}

func (b *LinkElement) String() string {
	return b.node.String()
}

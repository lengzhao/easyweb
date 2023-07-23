package e

import (
	"fmt"

	"github.com/lengzhao/easyweb/util"
)

type BoxElement struct {
	BaseElement
	node *HtmlNode
}

func Box(items ...any) *BoxElement {
	var out BoxElement
	out.id = util.GetID()
	out.node = NewNode("div")
	for _, v := range items {
		out.cont += fmt.Sprint(v)
	}
	return &out
}

func (b *BoxElement) Add(in ...any) *BoxElement {
	for _, v := range in {
		b.cont += fmt.Sprint(v)
	}
	return b
}

func (b *BoxElement) Class(in string) *BoxElement {
	if b.cls == "" {
		b.cls = in
	} else {
		b.cls += " " + in
	}
	return b
}

func (b *BoxElement) String() string {
	b.node.SetAttr("class", b.cls)
	b.node.SetHtml(b.cont)
	return b.node.String()
}

package e

import (
	"fmt"

	"github.com/lengzhao/easyweb/util"
)

type RowElement struct {
	BaseElement
}

func Row(childs ...any) *RowElement {
	var out RowElement
	out.id = util.GetID()
	out.cont = fmt.Sprint(childs...)
	return &out
}

func (b *RowElement) Write(in any) *RowElement {
	b.cont += fmt.Sprint(in)
	return b
}

func (b *RowElement) String() string {
	node := NewNode("div")
	if b.id != "" {
		node.SetAttr("id", b.id)
	}
	node.SetAttr("class", "row "+b.cls)
	node.SetText(b.cont)
	return node.String()
}

package e

import (
	"fmt"

	"github.com/lengzhao/easyweb/util"
)

type LabelElement struct {
	BaseElement
}

func Label(text string) *LabelElement {
	var out LabelElement
	out.cont = text
	out.id = util.GetID()
	return &out
}

func (b *LabelElement) Write(in any) *LabelElement {
	b.cont += fmt.Sprint(in)
	return b
}

func (b *LabelElement) String() string {
	node := NewNode("p")
	if b.id != "" {
		node.SetAttr("id", b.id)
	}
	node.SetText(b.cont)
	return node.String()
}

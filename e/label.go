package e

import (
	"fmt"
)

type LabelElement struct {
	BaseElement
}

func Label(text string) *LabelElement {
	var out LabelElement
	out.cont = text
	return &out
}

func (b *LabelElement) Write(in any) *LabelElement {
	b.cont += fmt.Sprint(in)
	return b
}

func (b *LabelElement) String() string {
	node := NewNode("p")
	node.SetText(b.cont)
	return node.String()
}

package e

import (
	"fmt"

	"github.com/lengzhao/easyweb/util"
)

type ListElement struct {
	BaseElement
	values []any
}

func List(in []any) *ListElement {
	var out ListElement
	out.id = util.GetID()
	out.values = append(out.values, in...)
	return &out
}

func (b *ListElement) Add(in []any) *ListElement {
	b.values = append(b.values, in...)
	return b
}

func (b *ListElement) String() string {
	node := NewNode("ul")
	node.AddAttribute("class", "list-group "+b.cls)
	for _, v := range b.values {
		node.AddChild(NewNode("li").AddAttribute("class", "list-group-item").SetHtml(fmt.Sprint(v)))
	}
	return node.String()
}

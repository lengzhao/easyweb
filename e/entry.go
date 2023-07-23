package e

import (
	"fmt"
)

type EntryElement struct {
	BaseElement
	name   string
	prefix string
	suffix string
	typ    string
}

func Entry(name string) *EntryElement {
	out := &EntryElement{}
	out.name = name
	out.typ = "text"

	return out
}

func (e *EntryElement) String() string {
	node := NewNode("div")
	node.SetAttr("class", "input-group "+e.cls)
	if e.id != "" {
		node.SetAttr("id", e.id)
	}
	if e.prefix != "" {
		node.AddChild(NewNode("span").SetAttr("class", "input-group-text").SetText(e.prefix))
	}
	btn := NewNode("input").SetAttr("type", e.typ).SetAttr("name", e.name).SetAttr("class", "form-control")
	btn.SetAttr("id", "file1")
	if e.cont != "" {
		btn.SetAttr("value", e.cont)
	}
	node.AddChild(btn)
	if e.suffix != "" {
		node.AddChild(NewNode("span").SetAttr("class", "input-group-text").SetText(e.suffix))
	}

	return node.String()
}

func (e *EntryElement) Prefix(in string) *EntryElement {
	e.prefix = in
	return e
}

func (e *EntryElement) SetType(in string) *EntryElement {
	e.typ = in
	return e
}

func (e *EntryElement) Suffix(in string) *EntryElement {
	e.suffix = in
	return e
}

func (e *EntryElement) Value(in any) *EntryElement {
	e.cont = fmt.Sprint(in)
	return e
}

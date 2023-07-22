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
	node.AddAttribute("class", "input-group "+e.cls)
	if e.id != "" {
		node.AddAttribute("id", e.id)
	}
	if e.prefix != "" {
		node.AddChild(NewNode("span").AddAttribute("class", "input-group-text").SetText(e.prefix))
	}
	btn := NewNode("input").AddAttribute("type", e.typ).AddAttribute("name", e.name).AddAttribute("class", "form-control")
	btn.AddAttribute("id", "file1")
	if e.cont != "" {
		btn.AddAttribute("value", e.cont)
	}
	node.AddChild(btn)
	if e.suffix != "" {
		node.AddChild(NewNode("span").AddAttribute("class", "input-group-text").SetText(e.suffix))
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

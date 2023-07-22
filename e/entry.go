package e

import (
	"fmt"
)

type EntryElement struct {
	BaseElement
	name   string
	prefix string
	suffix string
}

func Entry(name string) *EntryElement {
	out := &EntryElement{}
	out.name = name

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
	btn := NewNode("input").AddAttribute("type", "text").AddAttribute("name", e.name).AddAttribute("class", "form-control")
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

func (e *EntryElement) Suffix(in string) *EntryElement {
	e.suffix = in
	return e
}

func (e *EntryElement) Value(in any) *EntryElement {
	e.cont = fmt.Sprint(in)
	return e
}

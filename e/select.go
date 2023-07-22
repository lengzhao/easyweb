package e

import (
	"fmt"

	"github.com/lengzhao/easyweb/util"
)

type SelectElement struct {
	BaseElement
	name string
}

func Select(name string) *SelectElement {
	var out SelectElement
	if name == "" {
		name = util.GetCallerID(util.LevelParent)
	}
	out.name = name
	out.id = util.GetID()
	return &out
}

func (e *SelectElement) String() string {
	node := NewNode("select")
	if e.id != "" {
		node.AddAttribute("id", e.id)
	}
	node.AddAttribute("name", e.name)
	node.AddAttribute("class", "form-select "+e.cls)
	node.SetHtml(e.cont)

	return node.String()
}

func (e *SelectElement) Class(in string) *SelectElement {
	e.cls += " " + in
	return e
}

type SelectItem struct {
	Value    string
	Text     string
	Selected bool
}

func (e *SelectElement) Add(in any) *SelectElement {
	switch val := in.(type) {
	case map[string]string:
		for k, v := range val {
			e.cont += NewNode("option").AddAttribute("value", v).SetText(k).String()
		}
	case SelectItem:
		node := NewNode("option")
		node.SetText(val.Text)
		if val.Selected {
			node.AddAttribute("selected", "")
		}
		node.AddAttribute("value", val.Value)
		e.cont += node.String()

	case []SelectItem:
		for _, v := range val {
			node := NewNode("option")
			node.SetText(v.Text)
			if v.Selected {
				node.AddAttribute("selected", "")
			}
			node.AddAttribute("value", v.Value)
			e.cont += node.String()
		}
	default:
		e.cont += fmt.Sprint(in)
	}
	return e
}

package e

import (
	"fmt"

	"github.com/lengzhao/easyweb/util"
)

type TableElement struct {
	BaseElement
	name      string
	header    []string
	values    [][]any
	showIndex bool
}

func Table(name string) *TableElement {
	var out TableElement
	if name == "" {
		name = util.GetCallerID(util.LevelParent)
	}
	out.name = name
	out.id = util.GetID()
	return &out
}

func (e *TableElement) Class(in string) *TableElement {
	e.cls += " " + in
	return e
}
func (e *TableElement) Header(in []string) *TableElement {
	e.header = in
	return e
}

type TableItem struct {
	Value string
	Text  string
}

func (e *TableElement) SetValue(in [][]any) *TableElement {
	e.values = in
	return e
}

func (e *TableElement) ShowIndex() *TableElement {
	e.showIndex = true
	return e
}

func (e *TableElement) String() string {
	table := NewNode("table")
	if e.cls == "" {
		e.cls = "table-striped table-hover"
	}
	if e.id != "" {
		table.AddAttribute("id", e.id)
	}
	table.AddAttribute("class", "table "+e.cls)
	if e.name != "" {
		table.AddAttribute("id", e.name)
	}
	trh := NewNode("tr")
	if e.showIndex {
		trh.AddChild(NewNode("th").AddAttribute("scope", "col").SetText("#"))
	}
	for _, v := range e.header {
		trh.AddChild(NewNode("th").AddAttribute("scope", "col").SetText(v))
	}
	table.AddChild(NewNode("thead").AddChild(trh))
	tBody := NewNode("tbody")
	for i, it := range e.values {
		tr := NewNode("tr")
		if e.showIndex {
			tr.AddChild(NewNode("th").AddAttribute("scope", "row").SetText(fmt.Sprint(i + 1)))
		}
		for _, v := range it {
			tr.AddChild(NewNode("td").SetText(fmt.Sprint(v)))
		}
		tBody.AddChild(tr)
	}
	table.AddChild(tBody)

	return table.String()
}

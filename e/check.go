package e

import "fmt"

type CheckElement struct {
	BaseElement
	name    string
	checked bool
}

func Check(name string) *CheckElement {
	var out CheckElement
	out.name = name
	return &out
}

func (e *CheckElement) String() string {
	node := NewNode("div")
	node.SetAttr("class", "form-check "+e.cls)
	if e.id != "" {
		node.SetAttr("id", e.id)
	}
	c1 := NewNode("input")
	c1.SetAttr("type", "checkbox")
	c1.SetAttr("class", "form-check-input")
	c1.SetAttr("name", e.name)
	if e.checked {
		c1.SetAttr("checked", "")
	}
	node.AddChild(c1)
	c2 := NewNode("label")
	c2.SetAttr("class", "form-check-label")
	c2.SetText(e.cont)
	node.AddChild(c2)

	return node.String()
}

func (e *CheckElement) Class(in string) *CheckElement {
	e.cls += " " + in
	return e
}

func (e *CheckElement) SetChecked() *CheckElement {
	e.checked = true
	return e
}

func (e *CheckElement) Text(in any) *CheckElement {
	e.cont = fmt.Sprint(in)
	return e
}

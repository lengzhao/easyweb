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
	node.AddAttribute("class", "form-check "+e.cls)
	if e.id != "" {
		node.AddAttribute("id", e.id)
	}
	c1 := NewNode("input")
	c1.AddAttribute("type", "checkbox")
	c1.AddAttribute("class", "form-check-input")
	c1.AddAttribute("name", e.name)
	if e.checked {
		c1.AddAttribute("checked", "")
	}
	node.AddChild(c1)
	c2 := NewNode("label")
	c2.AddAttribute("class", "form-check-label")
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

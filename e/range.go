package e

import (
	"fmt"

	"github.com/lengzhao/easyweb/util"
)

type RangeInputElement struct {
	BaseElement
	name  string
	min   int
	max   int
	step  int
	value string
}

func RangeInput(name string) *RangeInputElement {
	var out RangeInputElement
	out.name = name
	out.max = 100
	out.id = util.GetID()
	return &out
}

func (b *RangeInputElement) SetRange(min, max, step int) *RangeInputElement {
	b.min = min
	b.max = max
	b.step = step
	return b
}

func (b *RangeInputElement) SetValue(value int) *RangeInputElement {
	b.value = fmt.Sprint(value)
	return b
}
func (b *RangeInputElement) Write(in any) *RangeInputElement {
	b.cont = fmt.Sprint(in)
	return b
}

func (b *RangeInputElement) String() string {
	//<label for="customRange3" class="form-label">Example range</label>
	//<input type="range" class="form-range" min="0" max="5" step="0.5" id="customRange3">
	cid := util.GetCallerID(util.LevelParent) + b.name
	base := NewNode("div")
	if b.id != "" {
		base.SetAttr("id", b.id)
	}
	node := NewNode("label")
	node.SetText(b.cont)
	node.SetAttr("class", "form-label")
	node.SetAttr("for", cid)
	base.AddChild(node)

	n2 := NewNode("input")
	n2.SetAttr("type", "range")
	n2.SetAttr("class", "form-range")
	if b.min > 0 {
		n2.SetAttr("min", fmt.Sprint(b.min))
	}
	if b.max > 0 {
		n2.SetAttr("max", fmt.Sprint(b.max))
	}
	if b.step > 0 {
		n2.SetAttr("step", fmt.Sprint(b.step))
	}
	n2.SetAttr("id", cid)
	if b.value != "" {
		n2.SetAttr("value", b.value)
	}
	base.AddChild(n2)
	return base.String()
}

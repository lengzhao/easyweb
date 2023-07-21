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
	node := NewNode("label")
	node.SetText(b.cont)
	node.AddAttribute("class", "form-label")
	node.AddAttribute("for", cid)

	n2 := NewNode("input")
	n2.AddAttribute("type", "range")
	n2.AddAttribute("class", "form-range")
	if b.min > 0 {
		n2.AddAttribute("min", fmt.Sprint(b.min))
	}
	if b.max > 0 {
		n2.AddAttribute("max", fmt.Sprint(b.max))
	}
	if b.step > 0 {
		n2.AddAttribute("step", fmt.Sprint(b.step))
	}
	n2.AddAttribute("id", cid)
	if b.value != "" {
		n2.AddAttribute("value", b.value)
	}
	return node.String() + n2.String()
}

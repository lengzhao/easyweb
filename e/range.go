package e

import (
	"fmt"
)

type rangeElement struct {
	HtmlToken
}

func RangeInput(name, text string) *rangeElement {
	var out rangeElement
	id := getID()
	out.parseText(`<div  class="input-group"><label for="` + id + `" class="input-group-text">` + text + `</label>
	<input type="range" class="form-range form-control" min="0" max="10" step="1" name="` + name + `" id="` + id + `" oninput="this.nextElementSibling.value = this.value"><output class="input-group-text"></output></div>`)
	out.SetAttr("id", id)
	return &out
}

func (b *rangeElement) SetRange(min, max, step int) *rangeElement {
	if min > max {
		max = min + 10
	}
	if step <= 0 {
		step = 1
	}
	b.children[1].SetAttr("min", fmt.Sprint(min))
	b.children[1].SetAttr("max", fmt.Sprint(max))
	b.children[1].SetAttr("step", fmt.Sprint(step))
	return b
}

func (b *rangeElement) SetValue(value int) *rangeElement {
	b.children[1].SetAttr("value", fmt.Sprint(value))
	b.children[2].text = fmt.Sprint(value)
	return b
}

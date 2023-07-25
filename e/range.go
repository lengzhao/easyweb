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
	out.parseText(`<div><label for="` + id + `" class="form-label">` + text + `</label>
	<input type="range" class="form-range" min="0" max="100" step="1" id="` + id + `"></div>`)
	out.Attr("id", id)
	return &out
}

func (b *rangeElement) SetRange(min, max, step int) *rangeElement {
	b.children[1].Attr("min", fmt.Sprint(min))
	b.children[1].Attr("max", fmt.Sprint(max))
	b.children[1].Attr("step", fmt.Sprint(step))
	return b
}

func (b *rangeElement) SetValue(value int) *rangeElement {
	b.children[1].Attr("value", fmt.Sprint(value))
	return b
}

package e

import "github.com/lengzhao/easyweb"

type datalistElement struct {
	HtmlToken
}

var _ IElement = &datalistElement{}

func Datalist(name, text string) *datalistElement {
	out := datalistElement{}
	did := getID()
	out.parseText(`<div class="input-group"><span class="input-group-text">` + text + `</span>
	<input class="form-control" list="` + did + `" name="` + name + `" required>
	<datalist id="` + did + `">
	</datalist></div>`)
	out.SetAttr("id", getID())
	return &out
}

func (e *datalistElement) AddValues(text ...string) *datalistElement {
	op, _ := ParseHtml(`<option value="">`)
	for _, it := range text {
		lop := op.Copy()
		lop.SetAttr("value", it)
		e.children[2].Add(lop)
	}
	return e
}

func (e *datalistElement) AddItem(value, text string) *datalistElement {
	op, _ := ParseHtml(`<option value="">`)
	op.SetAttr("value", value)
	op.SetAttr("label", text)
	e.children[2].Add(op)
	return e
}

func (e *datalistElement) SetChangeCb(cb func(p easyweb.Page, id string, value string)) *datalistElement {
	e.children[2].SetCb("change", func(p easyweb.Page, id string, dataType easyweb.CbDataType, data []byte) {
		cb(p, id, string(data))
	})
	return e
}

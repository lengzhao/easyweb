package e

import (
	"encoding/json"
	"fmt"
)

type FormElement struct {
	BaseElement
	cb func(id string, info map[string]string)
}

func Form() *FormElement {
	var out FormElement
	return &out
}

func (b *FormElement) GetType() string {
	return "form"
}

func (b *FormElement) String() string {
	var out string
	out += `<form class="` + b.cls + `">`
	out += b.cont
	out += `<button type="submit" class="btn btn-primary">Submit</button></form>`
	return out
}

type FormItem struct {
	Name  string
	Value string
	Text  string
}

func (b *FormElement) Add(in any) *FormElement {
	switch val := in.(type) {
	case map[string]string:
		for k, v := range val {
			b.cont += `<div class="input-group ">`
			b.cont += `<span class="input-group-text">` + v + `</span>`
			b.cont += `<input type="text" class="form-control" name="` + k + `"></div>`
		}
	case FormItem:
		b.cont += `<div class="input-group ">`
		b.cont += `<span class="input-group-text">` + val.Text + `</span>`
		b.cont += `<input type="text" class="form-control" name="` + val.Name + `" value="` + val.Value + `"></div>`
	case []FormItem:
		for _, v := range val {
			b.cont += `<div class="input-group ">`
			b.cont += `<span class="input-group-text">` + v.Text + `</span>`
			b.cont += `<input type="text" class="form-control" name="` + v.Name + `" value="` + v.Value + `"></div>`
		}
	default:
		b.cont += fmt.Sprint(in)
	}
	return b
}

func (b *FormElement) ElementCb(id, info string) {
	if b.cb != nil {
		var info2 map[string]string
		json.Unmarshal([]byte(info), &info2)
		b.cb(id, info2)
	}
}

func (b *FormElement) SetCb(cb func(id string, info map[string]string)) *FormElement {
	b.cb = cb
	return b
}

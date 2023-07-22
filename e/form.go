package e

import (
	"encoding/json"
	"fmt"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/util"
)

type FormElement struct {
	BaseElement
	action string
	method string
	cb     func(id string, info map[string]string)
}

var _ easyweb.Event = &FormElement{}

func Form(cb func(id string, info map[string]string)) *FormElement {
	var out FormElement
	out.cb = cb
	out.id = util.GetID()
	return &out
}

func (b *FormElement) Action(action string) *FormElement {
	b.action = action
	return b
}
func (b *FormElement) Method(method string) *FormElement {
	b.method = method
	return b
}

func (b *FormElement) String() string {
	node := NewNode("form")
	if b.id != "" {
		node.AddAttribute("id", b.id)
	}
	if b.action != "" {
		node.AddAttribute("action", b.action)
	}
	if b.method != "" {
		node.AddAttribute("method", b.method)
	}
	node.SetHtml(b.cont)
	//<button type="submit" class="btn btn-primary">Submit</button>
	node.AddChild(NewNode("button").AddAttribute("type", "submit").AddAttribute("class", "btn btn-primary").SetText("Submit"))
	return node.String()
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

func (b *FormElement) MessageCb(id, info string) {
	if b.cb != nil {
		var info2 map[string]string
		json.Unmarshal([]byte(info), &info2)
		b.cb(id, info2)
	}
}

func (b *FormElement) EventInfo() (string, string) {
	return b.id, string(EventForm)
}

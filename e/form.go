package e

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/util"
)

type FormElement struct {
	BaseElement
	action     string
	method     string
	enctype    string
	cb         func(id string, info map[string]string)
	fileCb     func(id, fn string, size int64, data []byte)
	closeEvent bool
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
	if b.action != "" {
		b.closeEvent = true
	} else {
		b.closeEvent = false
	}
	if b.enctype == "" {
		b.enctype = "multipart/form-data"
	}
	if b.method == "" {
		b.method = http.MethodPost
	}
	return b
}
func (b *FormElement) Method(method string) *FormElement {
	b.method = method
	return b
}

func (b *FormElement) SetFileCb(cb func(id, fn string, size int64, data []byte)) *FormElement {
	b.fileCb = cb
	return b
}

func (b *FormElement) String() string {
	node := NewNode("form")
	if b.id != "" {
		node.SetAttr("id", b.id)
	}
	if b.action != "" {
		node.SetAttr("action", b.action)
	}
	if b.method != "" {
		node.SetAttr("method", b.method)
	}
	if b.enctype != "" {
		node.SetAttr("enctype", b.enctype)
	}
	node.SetHtml(b.cont)
	//<button type="submit" class="btn btn-primary">Submit</button>
	node.AddChild(NewNode("button").SetAttr("type", "submit").SetAttr("class", "btn btn-primary").SetText("Submit"))
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

func (b *FormElement) FileCb(id, fn string, size int64, data []byte) {
	if b.fileCb != nil {
		b.fileCb(id, fn, size, data)
	}
}

func (b *FormElement) EventInfo() (string, string) {
	if b.closeEvent {
		return "", ""
	}
	return b.id, string(EventForm)
}

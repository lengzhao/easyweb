package e

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lengzhao/easyweb"
)

type formElement struct {
	HtmlToken
	cb     func(id string, info map[string]string)
	fileCb func(id string, data []byte)
}

var fromTempData string = `<form>
<div>
</div>
<button type="submit" class="btn btn-primary">Submit</button>
</form>`

func Form(cb func(id string, info map[string]string)) *formElement {
	var out formElement
	out.Parse(fromTempData)
	out.Attr("id", getID())
	out.cb = cb
	if cb != nil {
		out.SetCb(func(id string, data []byte) {
			if data[0] != byte(easyweb.CbDataTypeString) {
				if out.fileCb != nil {
					out.fileCb(id, data)
				}
				return
			}
			info := make(map[string]string)
			err := json.Unmarshal(data[1:], &info)
			if err != nil {
				return
			}
			cb(id, info)
		})
	}
	return &out
}

func (b *formElement) Action(action, enctype string) *formElement {
	b.Attr("action", action)
	if enctype == "" {
		enctype = "multipart/form-data"
	}
	b.Attr("enctype", enctype)
	b.Attr("method", http.MethodPost)
	return b
}

func (b *formElement) SetFileCb(cb func(id string, data []byte)) *formElement {
	b.fileCb = cb
	return b
}

type FormItem struct {
	Name  string
	Value string
	Text  string
}

func (b *formElement) Add(in any) *formElement {
	b.Traverse(func(ht *HtmlToken) error {
		if ht.info.Data != "div" {
			return nil
		}
		switch val := in.(type) {
		case map[string]string:
			for k, v := range val {
				ht.text += `<div class="input-group ">`
				ht.text += `<span class="input-group-text">` + v + `</span>`
				ht.text += `<input type="text" class="form-control" name="` + k + `"></div>`
			}
		case FormItem:
			ht.text += `<div class="input-group ">`
			ht.text += `<span class="input-group-text">` + val.Text + `</span>`
			ht.text += `<input type="text" class="form-control" name="` + val.Name + `" value="` + val.Value + `"></div>`
		case []FormItem:
			for _, v := range val {
				ht.text += `<div class="input-group ">`
				ht.text += `<span class="input-group-text">` + v.Text + `</span>`
				ht.text += `<input type="text" class="form-control" name="` + v.Name + `" value="` + v.Value + `"></div>`
			}
		case *HtmlToken:
			ht.add(in)
		default:
			ht.add(Box(in))
		}
		return fmt.Errorf("finish")
	})
	return b
}

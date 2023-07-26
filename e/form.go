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
	out.parseText(fromTempData)
	out.Attr("id", getID())
	out.cb = cb
	if cb != nil {
		out.SetCb("", func(id string, data []byte) {
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

func (b *formElement) AddInput(name, text string) *formElement {
	item := InputGroup(name, text)
	b.Traverse(func(ht *HtmlToken) error {
		if ht.info.Data != "div" {
			return nil
		}
		ht.add(item)
		return nil
	})
	return b
}

func (b *formElement) Add(in any) *formElement {
	b.Traverse(func(ht *HtmlToken) error {
		if ht.info.Data != "div" {
			return nil
		}
		switch val := in.(type) {
		case []iSelf:
			for _, v := range val {
				ht.add(v.Self())
			}
		case iSelf:
			ht.add(val.Self())
		default:
			ht.add(in)
		}
		return fmt.Errorf("finish")
	})
	return b
}

func (b *formElement) SetButtonText(text string) *formElement {
	b.Traverse(func(ht *HtmlToken) error {
		if ht.info.Data != "button" {
			return nil
		}
		if ht.GetAttr("type") != "submit" {
			return nil
		}
		ht.children = nil
		ht.text = text
		return fmt.Errorf("finish")
	})
	return b
}

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

func Form(cb func(id string, info map[string]string)) *formElement {
	var out formElement
	out.parseText(`<form>
		<div>
		</div>
		<div class="d-flex justify-content-between">
			<p></p>
			<button type="submit" class="btn btn-primary ml-auto">Submit</button>
		</div>
	</form>`)
	out.Attr("id", getID())
	out.cb = cb
	if cb != nil {
		out.SetCb("submit", out.eventCb)
	}
	return &out
}

func (b *formElement) eventCb(id string, dataType easyweb.CbDataType, data []byte) {
	if dataType == easyweb.CbDataTypeBinary {
		if b.fileCb != nil {
			b.fileCb(id, data)
		}
		return
	}
	if b.cb == nil {
		return
	}
	info := make(map[string]string)
	err := json.Unmarshal(data, &info)
	if err != nil {
		return
	}
	b.cb(id, info)
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
	if cb != nil {
		b.SetCb("submit", b.eventCb)
	}
	return b
}

func (b *formElement) AddInput(name, text string) *formElement {
	item := InputGroup(name, text)
	b.Traverse(func(parent string, ht *HtmlToken) error {
		if ht.Info.Data != "div" || parent != "form" {
			return nil
		}
		ht.add(item)
		return fmt.Errorf("finish")
	})
	return b
}

func (b *formElement) Add(in any) *formElement {
	b.Traverse(func(parent string, ht *HtmlToken) error {
		if ht.Info.Data != "div" || parent != "form" {
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
	b.Traverse(func(parent string, ht *HtmlToken) error {
		if ht.Info.Data != "button" {
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

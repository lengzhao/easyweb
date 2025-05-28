package e

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lengzhao/easyweb"
)

type formElement struct {
	HtmlToken
	cb           func(p easyweb.Page, id string, info map[string]string)
	fileCb       func(p easyweb.Page, id string, data []byte)
	resetAfterCb bool
}

var _ IElement = &formElement{}

func Form(cb func(p easyweb.Page, id string, info map[string]string)) *formElement {
	var out formElement
	out.parseText(`<form>
		<div>
		</div>
		<div class="d-flex justify-content-between">
			<p></p>
			<button type="submit" class="btn btn-primary ml-auto">Submit</button>
		</div>
	</form>`)
	out.SetAttr("id", getID())
	out.cb = cb
	if cb != nil {
		out.SetCb("submit", out.eventCb)
	}
	out.resetAfterCb = true
	return &out
}

func (b *formElement) eventCb(p easyweb.Page, id string, dataType easyweb.CbDataType, data []byte) {
	if dataType == easyweb.CbDataTypeBinary {
		if b.fileCb != nil {
			b.fileCb(p, id, data)
		}
		return
	}
	if b.cb == nil {
		return
	}
	if b.resetAfterCb {
		p.Replace(b)
	}
	info := make(map[string]string)
	err := json.Unmarshal(data, &info)
	if err != nil {
		return
	}
	b.cb(p, id, info)
}

func (b *formElement) Action(action, enctype string) *formElement {
	b.SetAttr("action", action)
	if enctype == "" {
		enctype = "multipart/form-data"
	}
	b.SetAttr("enctype", enctype)
	b.SetAttr("method", http.MethodPost)
	return b
}

func (b *formElement) SetFileCb(cb func(p easyweb.Page, id string, data []byte)) *formElement {
	b.fileCb = cb
	if cb != nil {
		b.SetCb("submit", b.eventCb)
	}
	return b
}

func (b *formElement) AddInput(name, text string) *formElement {
	item := InputGroup(name, text)
	b.Traverse(nil, func(parent, ht IElement) error {
		if parent == nil || ht.HtmlTag() != "div" || parent.HtmlTag() != "form" {
			return nil
		}
		ht.Add(&item.HtmlToken)
		return fmt.Errorf("finish")
	})
	return b
}

func (b *formElement) AddItem(in IElement) *formElement {
	b.Traverse(nil, func(parent, ht IElement) error {
		if parent == nil || ht.HtmlTag() != "div" || parent.HtmlTag() != "form" {
			return nil
		}
		ht.Add(in)
		return fmt.Errorf("finish")
	})
	return b
}

func (b *formElement) SetButtonText(text string) *formElement {
	b.Traverse(nil, func(parent, ht IElement) error {
		if ht.HtmlTag() != "button" {
			return nil
		}
		if ht.GetAttr("type") != "submit" {
			return nil
		}
		ht.SetChild()
		ht.SetText(text)
		return fmt.Errorf("finish")
	})
	return b
}

func (b *formElement) ResetAfterCb(on bool) *formElement {
	b.resetAfterCb = on
	return b
}

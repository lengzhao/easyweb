package e

import (
	"encoding/json"
	"fmt"

	"github.com/lengzhao/easyweb"
)

type modalElement struct {
	HtmlToken
	cb func(p easyweb.Session, id string, info map[string]string)
}

var _ IElement = &modalElement{}

func Modal(btnText, title string) *modalElement {
	out := modalElement{}
	id := getID()
	sid := getID()
	out.parseText(`<div>
	<button type="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#` + id + `">` + btnText + `</button>
	<div class="modal fade" tabindex="-1" id="` + id + `" aria-labelledby="` + sid + `" aria-hidden="true">
	<div class="modal-dialog">
	  <div class="modal-content">
		<div class="modal-header">
			<h5 class="modal-title" id="` + sid + `">` + title + `</h5>
			<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
		</div>
		<div class="modal-body">
		</div>
		<div class="modal-footer">
			<button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
		</div>
	  </div>
	</div>
  </div>
  </div>`)
	out.SetAttr("id", getID())

	return &out
}

func FormModal(btnText, title string, cb func(p easyweb.Session, id string, info map[string]string)) *modalElement {
	out := modalElement{}
	id := getID()
	sid := getID()
	out.parseText(`<div>
	<button type="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#` + id + `">` + btnText + `</button>
	<div class="modal fade" tabindex="-1" id="` + id + `" aria-labelledby="` + sid + `" aria-hidden="true">
	<div class="modal-dialog">
	  <div class="modal-content">
		<form>
			<div class="modal-header">
				<h5 class="modal-title" id="` + sid + `">` + title + `</h5>
				<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
			</div>
			<div class="modal-body">
			</div>
			<div class="modal-footer">
				<button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
				<button type="submit" class="btn btn-primary ml-auto" data-bs-dismiss="modal">Submit</button>
			</div>
	    </form>
	  </div>
	</div>
  </div>
  </div>`)
	out.SetAttr("id", getID())
	out.cb = cb
	if cb != nil {
		out.Traverse(nil, func(parent, ht IElement) error {
			if ht.HtmlTag() != "form" {
				return nil
			}
			ht.SetCb("submit", out.eventCb)
			return fmt.Errorf("finish")
		})
	}

	return &out
}

func (b *modalElement) eventCb(p easyweb.Session, id string, dataType easyweb.CbDataType, data []byte) {
	if b.cb == nil {
		return
	}
	p.Replace(b)
	info := make(map[string]string)
	err := json.Unmarshal(data, &info)
	if err != nil {
		return
	}
	b.cb(p, id, info)
}

func (e *modalElement) SetBody(body IElement) *modalElement {
	e.Traverse(nil, func(parent, ht IElement) error {
		if ht.HtmlTag() != "div" {
			return nil
		}
		if ht.GetAttr("class") != "modal-body" {
			return nil
		}
		ht.SetChild()
		ht.Add(body)
		return fmt.Errorf("finish")
	})
	return e
}

func (e *modalElement) AddBody(body IElement) *modalElement {
	e.Traverse(nil, func(parent, ht IElement) error {
		if ht.HtmlTag() != "div" {
			return nil
		}
		if ht.GetAttr("class") != "modal-body" {
			return nil
		}
		ht.Add(body)
		return fmt.Errorf("finish")
	})
	return e
}

func (e *modalElement) AddFooter(footer IElement) *modalElement {
	e.Traverse(nil, func(parent, ht IElement) error {
		if ht.HtmlTag() != "div" {
			return nil
		}
		if ht.GetAttr("class") != "modal-footer" {
			return nil
		}
		ht.Add(footer)
		return fmt.Errorf("finish")
	})
	return e
}

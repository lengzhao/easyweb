package e

import "fmt"

type modalElement struct {
	HtmlToken
}

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

func (e *modalElement) SetBody(body any) *modalElement {
	e.Traverse(func(parent string, token *HtmlToken) error {
		if token.Info.Data != "div" {
			return nil
		}
		if token.GetAttr("class") != "modal-body" {
			return nil
		}
		token.children = nil
		token.add(body)
		return fmt.Errorf("finish")
	})
	return e
}

func (e *modalElement) AddFooter(footer any) *modalElement {
	e.Traverse(func(parent string, token *HtmlToken) error {
		if token.Info.Data != "div" {
			return nil
		}
		if token.GetAttr("class") != "modal-footer" {
			return nil
		}
		token.add(footer)
		return fmt.Errorf("finish")
	})
	return e
}

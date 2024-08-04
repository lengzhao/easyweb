package e

import (
	"fmt"
	"log"

	"github.com/lengzhao/easyweb"
)

type paginationElement struct {
	HtmlToken
	pgCb func(p easyweb.Page, id, item string)
}

var _ IElement = &paginationElement{}

func Pagination(items []string, cb func(p easyweb.Page, id, item string)) *paginationElement {
	out := paginationElement{}
	out.parseText(`<nav aria-label="Page navigation example">
	<ul class="pagination justify-content-center">
	</ul>
  </nav>`)
	id := getID()
	out.SetAttr("id", id)
	out.Traverse(nil, func(parent, token IElement) error {
		if token.HtmlTag() != "ul" {
			return nil
		}
		for _, item := range items {
			it, _ := ParseHtml(`<li class="page-item"><a class="page-link" href="#">` + item + `</a></li>`)
			token.Add(it)
		}
		return fmt.Errorf("finish")
	})
	out.SetCb("click", out.eventCallback)
	out.pgCb = cb
	out.eventKey = id + " li"

	return &out
}

func (e *paginationElement) eventCallback(p easyweb.Page, id string, dataType easyweb.CbDataType, data []byte) {
	log.Println("eventCallback", id, dataType, string(data))
	e.pgCb(p, e.GetID(), string(data))
}

func (e *paginationElement) Active(item string) *paginationElement {
	finish := false
	e.Traverse(nil, func(parent, token IElement) error {
		if token.HtmlTag() != "li" {
			return nil
		}
		token.Traverse(nil, func(parent, lt IElement) error {
			if lt.GetText() == item {
				token.SetAttr("class", "page-item active")
				finish = true
				return fmt.Errorf("finish")
			}
			return nil
		})
		if finish {
			return fmt.Errorf("finish")
		}
		return nil
	})
	return e
}

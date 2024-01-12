package e

import (
	"fmt"
	"log"

	"github.com/lengzhao/easyweb"
)

type paginationElement struct {
	HtmlToken
	pgCb func(id, item string)
}

func Pagination(items []string, cb func(id, item string)) *paginationElement {
	out := paginationElement{}
	out.parseText(`<nav aria-label="Page navigation example">
	<ul class="pagination justify-content-center">
	</ul>
  </nav>`)
	id := getID()
	out.SetAttr("id", id)
	out.Traverse(func(parent string, token *HtmlToken) error {
		if token.Info.Data != "ul" {
			return nil
		}
		for _, item := range items {
			it, _ := ParseHtml(`<li class="page-item"><a class="page-link" href="#">` + item + `</a></li>`)
			token.add(it)
		}
		return fmt.Errorf("finish")
	})
	out.SetCb("click", out.eventCallback)
	out.pgCb = cb
	out.eventKey = id + " li"

	return &out
}

func (e *paginationElement) eventCallback(id string, dataType easyweb.CbDataType, data []byte) {
	log.Println("eventCallback", id, dataType, string(data))
	e.pgCb(e.GetID(), string(data))
}

func (e *paginationElement) Active(item string) *paginationElement {
	finish := false
	e.Traverse(func(parent string, token *HtmlToken) error {
		if token.Info.Data != "li" {
			return nil
		}
		token.Traverse(func(parent string, lt *HtmlToken) error {
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

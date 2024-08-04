package e

import "github.com/lengzhao/easyweb"

type buttonElement struct {
	HtmlToken
}

func Button(text string, cb func(p easyweb.Page, id string)) IElement {
	var out buttonElement
	out.parseText(`<button type="button" class="btn btn-primary">` + text + `</button>`)
	out.SetAttr("id", getID())
	if cb != nil {
		out.SetCb("click", func(p easyweb.Page, id string, dataType easyweb.CbDataType, data []byte) {
			cb(p, id)
		})
	}
	return &out
}

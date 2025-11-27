package e

import "github.com/lengzhao/easyweb"

func Button(text string, cb func(p easyweb.Session, id string)) IElement {
	out := MustParseHtml(`<button type="button" class="btn btn-primary">` + text + `</button>`)
	out.SetAttr("id", getID())
	if cb != nil {
		out.SetCb("click", func(s easyweb.Session, id string, dataType easyweb.CbDataType, data []byte) {
			cb(s, id)
		})
	}
	return out
}

package e

type buttonElement struct {
	HtmlToken
}

func Button(text string, cb func(id string)) *buttonElement {
	out := buttonElement{}
	out.parseText(`<button type="button" class="btn btn-primary">` + text + `</button>`)
	out.Attr("id", getID())
	if cb != nil {
		out.SetCb("click", func(id string, data []byte) {
			cb(id)
		})
	}
	return &out
}

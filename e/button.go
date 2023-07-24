package e

type buttonElement struct {
	HtmlToken
}

func Button(text string, cb func(id string)) *buttonElement {
	out := buttonElement{}
	out.Parse(`<button type="button" class="btn btn-primary">Button</button>`)
	out.Attr("id", getID())
	if cb != nil {
		out.SetCb(func(id string, data []byte) {
			cb(id)
		})
	}
	return &out
}

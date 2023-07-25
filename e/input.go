package e

type inputElement struct {
	HtmlToken
}

func InputGroup(name, text string) *inputElement {
	out := &inputElement{}
	out.parseText(`<div class="input-group">
	<span class="input-group-text">` + text + `</span>
	<input type="text" class="form-control" aria-label="" name="` + name + `"/>
	<span class="input-group-text"></span>
  </div>`)
	if text == "" {
		out.children[0].disable = true
	}
	out.children[2].disable = true

	return out
}

func (e *inputElement) Suffix(text string) *inputElement {
	e.children[2].text = text
	e.children[2].disable = false
	return e
}

func (e *inputElement) ChangeType(text string) *inputElement {
	e.children[1].Attr("type", text)
	return e
}

func (e *inputElement) Value(value string) *inputElement {
	e.children[1].Attr("value", value)
	return e
}

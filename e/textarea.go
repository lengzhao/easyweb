package e

import (
	"bytes"
	"fmt"
	"html/template"
)

type textareaElement struct {
	HtmlToken
}

func Textarea(name, title string) *textareaElement {
	var out textareaElement
	id := getID()
	out.parseText(`<div class="input-group" id="` + id + `">
	<span class="input-group-text">` + title + `</span>
	<textarea type="text" class="form-control" aria-label="` + name + `"> </textarea>
  </div>`)
	if name != "" {
		out.children[1].SetAttr("name", name)
	}
	return &out
}

func (e *textareaElement) Readonly(readonly bool) *textareaElement {
	if readonly {
		e.children[1].SetAttr("readonly", "true")
	} else {
		e.children[1].SetAttr("readonly", "")
	}
	return e
}

func (e *textareaElement) Rows(num uint) *textareaElement {
	e.children[1].SetAttr("rows", fmt.Sprintf("%d", num))
	return e
}

func (e *textareaElement) Set(text string) *textareaElement {
	buff := new(bytes.Buffer)
	template.HTMLEscape(buff, []byte(text))
	e.children[1].text = buff.String()
	return e
}

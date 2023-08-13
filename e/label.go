package e

import (
	"bytes"
	"fmt"
	"html/template"
)

type labelElement struct {
	HtmlToken
}

func Label(text string) *labelElement {
	out := labelElement{}
	out.parseText(`<label></label>`)
	out.Attr("id", getID())
	out.text = text
	return &out
}

func (e *labelElement) Set(text any) *labelElement {
	e.text = fmt.Sprint(text)
	buff := new(bytes.Buffer)
	template.HTMLEscape(buff, []byte(e.text))
	e.text = buff.String()
	return e
}

func (e *labelElement) Add(text any) *labelElement {
	e.text += fmt.Sprint(text)
	buff := new(bytes.Buffer)
	template.HTMLEscape(buff, []byte(e.text))
	e.text = buff.String()
	return e
}

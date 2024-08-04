package e

import "fmt"

type labelElement struct {
	HtmlToken
}

var _ IElement = &labelElement{}

func Label(text string) *labelElement {
	out := labelElement{}
	out.parseText(`<label></label>`)
	out.SetAttr("id", getID())
	out.text = text
	return &out
}

func (e *labelElement) AddText(text string) *labelElement {
	e.text += text
	return e
}

// <div>text</div>
func Div(text string) IElement {
	var out HtmlToken
	out.parseText(`<div>` + text + `</div>`)
	return &out
}

// <p>text</p>
func P(text string) IElement {
	var out HtmlToken
	out.parseText(`<p>` + text + `</p>`)
	return &out
}

func Text(text string) IElement {
	var out HtmlToken
	out.SetText(text)
	return &out
}

// h1-h6,level=1-6
func H(level uint, msg string) IElement {
	return MustParseHtml(fmt.Sprintf("<h%d>%s</h%d>", level, msg, level))
}

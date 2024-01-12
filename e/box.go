package e

type boxElement struct {
	HtmlToken
}

func Box(items ...any) *boxElement {
	out := boxElement{}
	out.parseText(`<div></div>`)
	out.SetAttr("id", getID())
	out.add(items...)
	return &out
}

func (e *boxElement) Add(items ...any) {
	e.add(items...)
}

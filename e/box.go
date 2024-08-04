package e

type boxElement struct {
	HtmlToken
}

func Box(items ...any) IElement {
	out := boxElement{}
	out.parseText(`<div></div>`)
	out.SetAttr("id", getID())
	out.AddAny(items...)
	return &out.HtmlToken
}

package e

type boxElement struct {
	HtmlToken
}

func Box(items ...any) *boxElement {
	out := boxElement{}
	out.Parse(`<div></div>`)
	out.Attr("id", getID())
	out.add(items...)
	return &out
}

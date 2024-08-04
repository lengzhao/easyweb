package e

type rowElement struct {
	HtmlToken
}

var _ IElement = &rowElement{}

func Row(items ...IElement) *rowElement {
	out := rowElement{}
	out.parseText(`<div class="row"></div>`)
	out.SetAttr("id", getID())
	out.Add(items...)
	return &out
}

func (e *rowElement) AddCols(items ...IElement) *rowElement {
	for _, it := range items {
		e.Add(Col(0, it))
	}
	return e
}

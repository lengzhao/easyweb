package e

type rowElement struct {
	HtmlToken
}

func Row(items ...any) *rowElement {
	out := rowElement{}
	out.parseText(`<div class="row"></div>`)
	out.SetAttr("id", getID())
	out.add(items...)
	return &out
}

func (e *rowElement) Add(items ...any) *rowElement {
	e.add(items...)
	return e
}

func (e *rowElement) AddCols(items ...any) *rowElement {
	for _, it := range items {
		e.add(Col(0, it))
	}
	return e
}

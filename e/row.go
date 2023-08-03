package e

type rowElement struct {
	HtmlToken
}

func Row(items ...any) *rowElement {
	out := rowElement{}
	out.parseText(`<div class="row"></div>`)
	out.Attr("id", getID())
	out.add(items...)
	return &out
}

func (e *rowElement) Add(items ...any) {
	e.add(items...)
}

func (e *rowElement) AddCols(items ...any) {
	for _, it := range items {
		e.add(Col(0, it))
	}
}

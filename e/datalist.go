package e

type datalistElement struct {
	HtmlToken
}

func Datalist(name, text string) *datalistElement {
	out := datalistElement{}
	did := getID()
	out.parseText(`<div class="input-group"><span class="input-group-text">` + text + `</span>
	<input class="form-control" list="` + did + `" name="` + name + `" required>
	<datalist id="` + did + `">
	</datalist></div>`)
	out.SetAttr("id", getID())
	return &out
}

func (e *datalistElement) Add(text ...string) *datalistElement {
	op, _ := ParseHtml(`<option value="">`)
	for _, it := range text {
		lop := op.Copy()
		lop.SetAttr("value", it)
		e.children[2].Add(lop)
	}
	return e
}

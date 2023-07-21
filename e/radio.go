package e

import "fmt"

type RadioElement struct {
	BaseElement
	name string
}

func Radio(name string) *RadioElement {
	var out RadioElement
	out.name = name
	return &out
}

func (e *RadioElement) String() string {

	return e.cont
}

func (e *RadioElement) Class(in string) *RadioElement {
	e.cls += " " + in
	return e
}

type RadioItem struct {
	Value string
	Text  string
}

func (e *RadioElement) Add(in any) *RadioElement {
	switch val := in.(type) {
	case map[string]string:
		for k, v := range val {
			lid := fmt.Sprintf("%s%04d", e.name, len(e.cont))
			e.cont += `<div class="form-check ` + e.cls + `">`
			e.cont += `<input class="form-check-input" type="radio" name="` + e.name + `" value="` + v + `" id="` + lid + `">`
			e.cont += `<label class="form-check-label" for="` + lid + `">` + k + `</label>`
			e.cont += `</div>`
		}
	case []RadioItem:
		for _, v := range val {
			lid := fmt.Sprintf("%s%04d", e.name, len(e.cont))
			e.cont += `<div class="form-check ` + e.cls + `">`
			e.cont += `<input class="form-check-input" type="radio" name="` + e.name + `" value="` + v.Value + `" id="` + lid + `">`
			e.cont += `<label class="form-check-label" for="` + lid + `>` + v.Text + `</label>`
			e.cont += `</div>`
		}
	case RadioItem:
		lid := fmt.Sprintf("%s%04d", e.name, len(e.cont))
		e.cont += `<div class="form-check ` + e.cls + `">`
		e.cont += `<input class="form-check-input" type="radio" name="` + e.name + `" value="` + val.Value + `" id="` + lid + `">`
		e.cont += `<label class="form-check-label" for="` + lid + `>` + val.Text + `</label>`
		e.cont += `</div>`
	default:
		e.cont += fmt.Sprint(in)
	}
	return e
}

package e

import "fmt"

type SelectElement struct {
	BaseElement
	name string
}

func Select(name string) *SelectElement {
	var out SelectElement
	out.name = name
	return &out
}

func (e *SelectElement) String() string {
	out := `<select class="form-select ` + e.cls + `" name="` + e.name + `">`
	out += e.cont
	out += `</select>`

	return out
}

func (e *SelectElement) Class(in string) *SelectElement {
	e.cls += " " + in
	return e
}

type SelectItem struct {
	Value    string
	Text     string
	Selected bool
}

func (e *SelectElement) Add(in any) *SelectElement {
	switch val := in.(type) {
	case map[string]string:
		for k, v := range val {
			e.cont += `<option value="` + v + `">` + k + `</option>`
		}
	case SelectItem:
		e.cont += `<option value="` + val.Value + `"`
		if val.Selected {
			e.cont += " selected"
		}
		e.cont += `>` + val.Text + `</option>`
	case []SelectItem:
		for _, v := range val {
			e.cont += `<option value="` + v.Value + `"`
			if v.Selected {
				e.cont += " selected"
			}
			e.cont += `>` + v.Text + `</option>`
		}
	default:
		e.cont += fmt.Sprint(in)
	}
	return e
}

package e

import "fmt"

type EntryElement struct {
	BaseElement
	name   string
	prefix string
	suffix string
}

func Entry(name string) *EntryElement {
	out := &EntryElement{}
	out.name = name

	return out
}

func (e *EntryElement) String() string {
	out := `<div class="input-group ` + e.cls + `">`
	if e.prefix != "" {
		out += `<span class="input-group-text">` + e.prefix + `</span>`
	}

	out += `<input type="text" class="form-control" value="` + e.cont + `" name="` + e.name + `">`
	if e.suffix != "" {
		out += `<span class="input-group-text">` + e.suffix + `</span>`
	}
	out += `</div>`

	return out
}

func (e *EntryElement) Prefix(in string) *EntryElement {
	e.prefix = in
	return e
}

func (e *EntryElement) Suffix(in string) *EntryElement {
	e.suffix = in
	return e
}

func (e *EntryElement) Value(in any) *EntryElement {
	e.cont = fmt.Sprint(in)
	return e
}

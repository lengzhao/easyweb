package e

import "github.com/lengzhao/easyweb"

type InputType string

const (
	InputTypeButton   InputType = "button"
	InputTypeCheckbox InputType = "checkbox"
	InputTypeColor    InputType = "color"
	InputTypeDate     InputType = "date"
	InputTypeDatetime InputType = "datetime-local"
	InputTypeEmail    InputType = "email"
	InputTypeFile     InputType = "file"
	InputTypeHidden   InputType = "hidden"
	InputTypeImage    InputType = "image"
	InputTypeMonth    InputType = "month"
	InputTypeNumber   InputType = "number"
	InputTypePassword InputType = "password"
	InputTypeRadio    InputType = "radio"
	InputTypeRange    InputType = "range"
	InputTypeReset    InputType = "reset"
	InputTypeSearch   InputType = "search"
	InputTypeSubmit   InputType = "submit"
	InputTypeTel      InputType = "tel"
	InputTypeText     InputType = "text"
	InputTypeTime     InputType = "time"
	InputTypeUrl      InputType = "url"
	InputTypeWeek     InputType = "week"
)

type inputElement struct {
	HtmlToken
}

func InputGroup(name, text string) *inputElement {
	out := &inputElement{}
	out.parseText(`<div class="input-group">
	<span class="input-group-text">` + text + `</span>
	<input type="text" class="form-control" aria-label="` + name + `" required/>
	<span class="input-group-text"></span>
  </div>`)
	if text == "" {
		out.children[0].SetAttr("hidden", "true")
	}
	out.children[2].SetAttr("hidden", "true")
	if name != "" {
		out.children[1].SetAttr("name", name)
	}
	out.SetAttr("id", getID())

	return out
}

func (e *inputElement) Suffix(text string) *inputElement {
	e.children[2].SetText(text)
	e.children[2].SetAttr("hidden", "")
	return e
}

func (e *inputElement) ChangeType(text InputType) *inputElement {
	e.children[1].SetAttr("type", string(text))
	if text == InputTypeNumber {
		e.children[1].SetAttr("step", "any")
		e.children[1].SetAttr("inputmode", "decimal")
	}
	return e
}

func (e *inputElement) Value(value string) *inputElement {
	e.children[1].SetAttr("value", value)
	return e
}

func (e *inputElement) ChangeSuffix(suffix IElement) *inputElement {
	e.children[2] = suffix
	return e
}

func (e *inputElement) ChangeInput(in IElement) *inputElement {
	e.children[1] = in
	return e
}

func (e *inputElement) Unrequired() *inputElement {
	e.children[1].SetAttr("required", "")
	return e
}

func (e *inputElement) Hidden() *inputElement {
	e.SetAttr("hidden", "true")
	return e
}

func (e *inputElement) SetChangeCb(cb func(p easyweb.Page, id string, value string)) *inputElement {
	e.children[1].SetCb("change", func(p easyweb.Page, id string, dataType easyweb.CbDataType, data []byte) {
		cb(p, id, string(data))
	})
	return e
}

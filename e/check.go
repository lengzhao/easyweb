package e

type checkElement struct {
	HtmlToken
}

func Check(name, text string) *checkElement {
	var out checkElement
	id := getID()
	out.parseText(`<div class="form-check">
	<input class="form-check-input" type="checkbox" value="" name="` + name + `" id="` + id + `"/>
	<label class="form-check-label" for="` + id + `">
	  ` + text + `
	</label>
  </div>`)
	return &out
}

func (e *checkElement) SetChecked() *checkElement {
	e.children[0].SetAttr("checked", "true")
	return e
}

package e

type accordionElement struct {
	HtmlToken
}

func Accordion() *accordionElement {
	var out accordionElement
	out.parseText(`<div class="accordion"></div>`)
	out.SetAttr("id", getID())
	return &out
}

func (e *accordionElement) AddItem(header, text string) *accordionElement {
	var Name string = e.GetAttr("id")
	var ItemID string = getID()
	if len(e.children) > 0 {
		tempText := `<div class="accordion-item">
						<h2 class="accordion-header">
							<button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#` + ItemID + `" aria-expanded="false" aria-controls="` + ItemID + `">
							` + header + `
							</button>
						</h2>
						<div id="` + ItemID + `" class="accordion-collapse collapse" data-bs-parent="#` + Name + `">
						<div class="accordion-body">
							` + text + `
						</div>
						</div>
					</div>`
		item, _ := ParseHtml(tempText)
		e.Add(item)
	} else {
		tempText := `<div class="accordion-item">
						<h2 class="accordion-header">
							<button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#` + ItemID + `" aria-expanded="true" aria-controls="` + ItemID + `">
							` + header + `
							</button>
						</h2>
						<div id="` + ItemID + `" class="accordion-collapse collapse show" data-bs-parent="#` + Name + `">
						<div class="accordion-body">
							` + text + `
						</div>
						</div>
					</div>`
		item, _ := ParseHtml(tempText)
		e.Add(item)

	}

	return e
}

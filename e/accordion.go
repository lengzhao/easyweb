package e

type accordionElement struct {
	HtmlToken
}

func Accordion() *accordionElement {
	var out accordionElement
	out.Parse(`<div class="accordion"></div>`)
	out.Attr("id", getID())
	return &out
}

func (e *accordionElement) AddItem(header, text string) *accordionElement {
	var Name string = e.GetAttr("id")
	var ItemID string = getID()
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
	item := parseStringToToken(tempText)
	if len(e.children) == 0 {
		item.children[0].children[0].Attr("aria-expanded", "true").Attr("class", "accordion-button")
		item.children[1].Attr("class", "accordion-collapse collapse show")
	}
	e.add(item)

	return e
}

type AccordionItem struct {
	Header string
	Text   string
}

func (e *accordionElement) Add(in any) *accordionElement {
	switch val := in.(type) {
	case AccordionItem:
		e.AddItem(val.Header, val.Text)
	case []AccordionItem:
		for _, v := range val {
			e.AddItem(v.Header, v.Text)
		}
	default:
		e.add(Box(in))
	}
	return e
}

package e

type selectElement struct {
	HtmlToken
}

func Select(name string) *selectElement {
	var out selectElement
	out.parseText(`<select class="form-select" aria-label="Default select"></select>`)
	out.SetAttr("id", getID())
	out.SetAttr("name", name)
	return &out
}

func (e *selectElement) Add(value, text string) *selectElement {
	item, _ := ParseHtml(`<option value="` + value + `">` + text + `</option>`)
	if len(e.children) == 0 {
		item.SetAttr("selected", "true")
	}
	e.add(item)

	return e
}

func (e *selectElement) Select(value string) *selectElement {
	// fmt.Println("set select:", value)
	e.Traverse(func(parent string, ht *HtmlToken) error {
		if ht.Info.Data != "option" || parent != "select" {
			return nil
		}
		if ht.GetAttr("value") == value {
			// fmt.Println("success to set:", ht.String())
			ht.SetAttr("selected", "true")
		} else {
			ht.SetAttr("selected", "")
		}
		return nil
	})
	return e
}

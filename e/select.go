package e

import "fmt"

type selectElement struct {
	HtmlToken
}

var _ IElement = &selectElement{}

func Select(name string) *selectElement {
	var out selectElement
	out.parseText(`<select class="form-select" aria-label="Default select"></select>`)
	out.SetAttr("id", getID())
	out.SetAttr("name", name)
	return &out
}

func (e *selectElement) SetItems(items ...string) *selectElement {
	e.children = nil
	for _, item := range items {
		e.AddItem(item, item)
	}

	return e
}

func (e *selectElement) AddItem(value, text string) *selectElement {
	item, _ := ParseHtml(`<option value="` + value + `">` + text + `</option>`)
	if len(e.children) == 0 {
		item.SetAttr("selected", "true")
	}
	e.Add(item)

	return e
}

func (e *selectElement) Select(value string) *selectElement {
	// fmt.Println("set select:", value)
	e.Traverse(nil, func(parent, ht IElement) error {
		if parent == nil || ht.HtmlTag() != "option" || parent.HtmlTag() != "select" {
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

func (e *selectElement) SetMultiple(multiple bool) *selectElement {
	if multiple {
		e.SetAttr("multiple", "true")
	} else {
		e.SetAttr("multiple", "")
	}
	return e
}

// show size(rows)
func (e *selectElement) SetSize(size uint) *selectElement {
	if size > 0 {
		e.SetAttr("size", fmt.Sprintf("%d", size))
	} else {
		e.SetAttr("size", "")
	}
	return e
}

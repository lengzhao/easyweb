package e

import "github.com/lengzhao/easyweb"

type dropdownElement struct {
	HtmlToken
}

var _ IElement = &dropdownElement{}

func Dropdown(text string) *dropdownElement {
	var out dropdownElement
	id := getID()
	out.parseText(`<div class="dropdown">
	<button class="btn btn-secondary dropdown-toggle" type="button" id="` + id + `" data-bs-toggle="dropdown" aria-expanded="false">
	  ` + text + `
	</button>
	<ul class="dropdown-menu" aria-labelledby="` + id + `">
	</ul>
  </div>`)
	out.SetAttr("id", getID())
	return &out
}

func (b *dropdownElement) AddItem(in ...IElement) *dropdownElement {
	b.Traverse(nil, func(parent, ht IElement) error {
		if ht.HtmlTag() != "ul" {
			return nil
		}
		for _, v := range in {
			item, _ := ParseHtml(`<li class="dropdown-item"></li>`)
			item.Add(v)
			ht.Add(item)
		}
		return nil
	})
	return b
}

func (b *dropdownElement) AddLink(text, url string) *dropdownElement {
	b.Traverse(nil, func(parent, ht IElement) error {
		if ht.HtmlTag() != "ul" {
			return nil
		}
		item, _ := ParseHtml(`<li><a class="dropdown-item" href="` + url + `">` + text + `</a></li>`)
		ht.Add(item)
		return nil
	})
	return b
}
func (b *dropdownElement) AddButton(text string, cb func(p easyweb.Page, id string)) *dropdownElement {
	b.Traverse(nil, func(parent, ht IElement) error {
		if ht.HtmlTag() != "ul" {
			return nil
		}
		btn := Button(text, cb)
		btn.SetAttr("class", "dropdown-item")
		item, _ := ParseHtml(`<li></li>`)
		item.Add(btn)
		ht.Add(item)
		return nil
	})
	return b
}

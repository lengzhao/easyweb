package e

type dropdownElement struct {
	HtmlToken
}

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
	return &out
}

func (b *dropdownElement) Add(in ...any) *dropdownElement {
	b.Traverse(func(parent string, ht *HtmlToken) error {
		if ht.Info.Data != "ul" {
			return nil
		}
		for _, v := range in {
			item, _ := ParseHtml(`<li class="dropdown-item"></li>`)
			item.add(v)
			ht.add(item)
		}
		return nil
	})
	return b
}

func (b *dropdownElement) AddLink(text, url string) *dropdownElement {
	b.Traverse(func(parent string, ht *HtmlToken) error {
		if ht.Info.Data != "ul" {
			return nil
		}
		item, _ := ParseHtml(`<li><a class="dropdown-item" href="` + url + `">` + text + `</a></li>`)
		ht.add(item)
		return nil
	})
	return b
}
func (b *dropdownElement) AddButton(text string, cb func(id string)) *dropdownElement {
	b.Traverse(func(parent string, ht *HtmlToken) error {
		if ht.Info.Data != "ul" {
			return nil
		}
		btn := Button(text, cb)
		btn.SetAttr("class", "dropdown-item")
		item, _ := ParseHtml(`<li></li>`)
		item.add(btn)
		ht.add(item)
		return nil
	})
	return b
}

package e

import (
	"fmt"
)

/*
<nav>
  <div class="nav nav-tabs" id="nav-tab" role="tablist">
    <button class="nav-link active" id="nav-home-tab" data-bs-toggle="tab" data-bs-target="#nav-home" type="button" role="tab" aria-controls="nav-home" aria-selected="true">Home</button>
    <button class="nav-link" id="nav-profile-tab" data-bs-toggle="tab" data-bs-target="#nav-profile" type="button" role="tab" aria-controls="nav-profile" aria-selected="false">Profile</button>
    <button class="nav-link" id="nav-contact-tab" data-bs-toggle="tab" data-bs-target="#nav-contact" type="button" role="tab" aria-controls="nav-contact" aria-selected="false">Contact</button>
  </div>
</nav>
<div class="tab-content" id="nav-tabContent">
  <div class="tab-pane fade show active" id="nav-home" role="tabpanel" aria-labelledby="nav-home-tab">...</div>
  <div class="tab-pane fade" id="nav-profile" role="tabpanel" aria-labelledby="nav-profile-tab">...</div>
  <div class="tab-pane fade" id="nav-contact" role="tabpanel" aria-labelledby="nav-contact-tab">...</div>
</div>
*/

type tabsElement struct {
	HtmlToken
}

var _ IElement = &tabsElement{}

func Tabs() *tabsElement {
	var out tabsElement
	out.parseText(`<div><nav>
			<div class="nav nav-tabs" role="tablist">
			</div></nav>
		<div class="tab-content">
		</div></div>`)
	out.SetAttr("id", getID())

	return &out
}

func (b *tabsElement) AddItem(title string, body IElement) *tabsElement {
	hid := getID()
	id := getID()
	hd, _ := ParseHtml(`<button class="nav-link" id="` + hid + `" data-bs-toggle="tab" data-bs-target="#` + id + `" type="button" role="tab" aria-controls="` + id + `" aria-selected="true">` + title + `</button>`)
	bd, _ := ParseHtml(`<div class="tab-pane fade" id="` + id + `" role="tabpanel" aria-labelledby="` + hid + `"></div>`)
	bd.Add(body)
	b.Traverse(nil, func(parent, ht IElement) error {
		if ht.GetAttr("role") == "tablist" {
			if len(ht.GetChilds()) == 0 {
				hd.SetAttr("class", "nav-link active")
			}
			ht.Add(hd)
		}
		if ht.GetAttr("class") == "tab-content" {
			if len(ht.GetChilds()) == 0 {
				bd.SetAttr("class", "tab-pane fade show active")
			}
			ht.Add(bd)
			return fmt.Errorf("finish")
		}
		return nil
	})
	return b
}

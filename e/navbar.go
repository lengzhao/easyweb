package e

import (
	"encoding/json"
	"fmt"

	"github.com/lengzhao/easyweb"
)

type navbarElement struct {
	HtmlToken
}

// Navbar creates a new navbarElement with the given name.
//
// Parameters:
// - name: a string representing the name of the navbarElement.
//
// Returns:
// - a pointer to the created navbarElement.
func Navbar(name string) *navbarElement {
	var out navbarElement
	sid := "navbar-" + getID()
	out.parseText(`<nav class="navbar navbar-expand-lg navbar-light bg-light">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">` + name + `</a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#` + sid + `" aria-controls="` + sid + `" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="` + sid + `">
        <ul class="navbar-nav me-auto my-2 my-lg-0 navbar-nav-scroll" style="--bs-scroll-height: 100px;">
        </ul>
		<form class="d-flex" hidden>
			<input class="form-control me-2" type="search" placeholder="Search" name="" aria-label="Search"/>
			<button class="btn btn-outline-success" type="submit">Search</button>
		</form>
      </div>
    </div>
  </nav>`)
	out.SetAttr("id", getID())
	return &out
}

func (b *navbarElement) AddItem(item ...IElement) *navbarElement {
	b.Traverse(nil, func(parent, ht IElement) error {
		if parent == nil || ht.HtmlToken().Data != "ul" || parent.HtmlToken().Data != "div" {
			return nil
		}
		for _, it := range item {
			lit, _ := ParseHtml(`<li class="nav-item"></li>`)
			if len(ht.GetChilds()) == 0 {
				it.SetAttr("class", it.GetAttr("class")+" nav-link active")
			} else {
				it.SetAttr("class", it.GetAttr("class")+" nav-link")
			}
			lit.Add(it)
			ht.Add(lit)
		}
		return fmt.Errorf("finish")
	})

	return b
}

func (b *navbarElement) Add(text, url string) *navbarElement {
	it := Link(text, url)
	b.AddItem(&it.HtmlToken)
	return b
}

func (b *navbarElement) SetSearchCb(cb func(p easyweb.Page, value string)) *navbarElement {
	id := getID()
	b.Traverse(nil, func(parent, ht IElement) error {
		if ht.HtmlToken().Data != "form" {
			return nil
		}
		ht.SetAttr("id", id)
		child := ht.GetChilds()
		if len(child) > 0 {
			child[0].SetAttr("name", id)
		}
		ht.SetAttr("hidden", "")
		ht.SetCb("submit", func(p easyweb.Page, id string, dataType easyweb.CbDataType, data []byte) {
			info := make(map[string]string)
			json.Unmarshal(data, &info)
			cb(p, info[id])
		})
		return fmt.Errorf("finish")
	})

	return b
}

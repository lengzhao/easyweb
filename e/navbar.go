package e

import (
	"encoding/json"
	"fmt"
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
		<form class="d-flex">
			<input class="form-control me-2" type="search" placeholder="Search" name="" aria-label="Search"/>
			<button class="btn btn-outline-success" type="submit">Search</button>
		</form>
      </div>
    </div>
  </nav>`)
	out.Attr("id", getID())
	out.Traverse(func(ht *HtmlToken) error {
		if ht.info.Data == "form" {
			ht.disable = true
			return fmt.Errorf("finish")
		}
		return nil
	})
	return &out
}

func (b *navbarElement) AddItem(item ...*HtmlToken) *navbarElement {
	b.Traverse(func(ht *HtmlToken) error {
		if ht.info.Data != "ul" {
			return nil
		}
		for _, it := range item {
			lit, _ := ParseHtml(`<li class="nav-item"></li>`)
			if len(ht.children) == 0 {
				it.Attr("class", it.GetAttr("class")+" nav-link active")
			} else {
				it.Attr("class", it.GetAttr("class")+" nav-link")
			}
			lit.add(it)
			ht.add(lit)
		}
		return fmt.Errorf("finish")
	})

	return b
}

func (b *navbarElement) Add(text, url string) *navbarElement {
	it := Link(text, url)
	b.AddItem(it.Self())
	return b
}

func (b *navbarElement) AddSearchItem(name string, cb func(value string)) *navbarElement {
	id := getID()
	b.Traverse(func(ht *HtmlToken) error {
		if ht.info.Data != "form" {
			return nil
		}
		if !ht.disable {
			return nil
		}
		ht.Attr("id", id)
		ht.children[0].Attr("name", name)
		ht.disable = false
		ht.SetCb("form", func(id string, data []byte) {
			info := make(map[string]string)
			json.Unmarshal(data[1:], &info)
			cb(info[name])
		})
		return fmt.Errorf("finish")
	})

	return b
}

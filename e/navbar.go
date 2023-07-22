package e

import (
	"fmt"

	"github.com/lengzhao/easyweb/util"
)

type NavbarElement struct {
	BaseElement
	name string
}

// Navbar creates a new NavbarElement with the given name.
//
// Parameters:
// - name: a string representing the name of the NavbarElement.
//
// Returns:
// - a pointer to the created NavbarElement.
func Navbar(name string) *NavbarElement {
	var out NavbarElement
	if name == "" {
		name = util.GetCallerID(util.LevelParent)
	}
	out.name = name
	return &out
}

func (b *NavbarElement) String() string {
	out := `<nav class="navbar navbar-expand-lg bg-body-tertiary">
	<div class="container-fluid">
	  <a class="navbar-brand" href="#">` + b.name + `</a>
	  <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
		<span class="navbar-toggler-icon"></span>
	  </button>
	  <div class="collapse navbar-collapse" id="navbarNav">
		<ul class="navbar-nav">
		  ` + b.cont + `
		</ul>
	  </div>
	</div>
  </nav>`
	return out
}

type NavbarItem struct {
	Name string
	Link string
}

func (b *NavbarElement) Write(in any) *NavbarElement {
	switch val := in.(type) {
	case map[string]string:
		for k, v := range val {
			b.cont += `<li class="nav-item"><a class="nav-link" href="` + v + `">` + k + `</a></li>`
		}
	case NavbarItem:
		b.cont += `<li class="nav-item"><a class="nav-link" href="` + val.Link + `">` + val.Name + `</a></li>`
	case []NavbarItem:
		for _, v := range val {
			b.cont += `<li class="nav-item"><a class="nav-link" href="` + v.Link + `">` + v.Name + `</a></li>`
		}
	default:
		b.cont += fmt.Sprint(in)
	}
	return b
}

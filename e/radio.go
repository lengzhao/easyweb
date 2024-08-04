package e

import (
	"strings"

	"github.com/lengzhao/easyweb/util"
)

type radioElement struct {
	HtmlToken
	name string
}

func Radio(name string) *radioElement {
	var out radioElement
	out.parseText("<div></div>")
	if name == "" {
		name = util.GetCallerID(util.LevelParent)
	}
	out.name = name
	out.SetAttr("id", getID())
	return &out
}

func (e *radioElement) AddItem(value, text string) *radioElement {
	id := getID()
	if len(e.children) == 0 {
		item, _ := ParseHtml(`<div class="form-check">
		<input class="form-check-input" type="radio" name="` + e.name + `" id="` + id + `" value="` + value + `" checked/>
		<label class="form-check-label" for="` + id + `">` + text + `</label>
	  </div>`)
		e.Add(item)
	} else {
		item, _ := ParseHtml(`<div class="form-check">
		<input class="form-check-input" type="radio" name="` + e.name + `" id="` + id + `" value="` + value + `"/>
		<label class="form-check-label" for="` + id + `">` + text + `</label>
	  </div>`)
		e.Add(item)
	}

	return e
}

func (e *radioElement) Select(value string) *radioElement {
	// fmt.Println("set check:", value)
	e.Traverse(nil, func(parent, ht IElement) error {
		if ht.HtmlTag() != "input" {
			return nil
		}
		if ht.GetAttr("value") == value {
			// fmt.Println("success to set:", ht.String())
			ht.SetAttr("checked", "true")
		} else {
			ht.SetAttr("checked", "")
		}
		return nil
	})
	return e
}

// Inline Make items appear inline. Must be called after adding items.
func (e *radioElement) Inline() *radioElement {
	e.Traverse(nil, func(parent, ht IElement) error {
		if parent == nil || ht.HtmlTag() != "div" || parent.HtmlTag() != "div" {
			return nil
		}
		cls := ht.GetAttr("class")
		if !strings.Contains(cls, "form-check") {
			return nil
		}
		ht.SetAttr("class", "form-check form-check-inline")
		return nil
	})
	return e
}

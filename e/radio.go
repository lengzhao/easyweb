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
	out.Attr("id", getID())
	return &out
}

func (e *radioElement) Add(value, text string) *radioElement {
	id := getID()
	item, _ := ParseHtml(`<div class="form-check">
	<input class="form-check-input" type="radio" name="` + e.name + `" id="` + id + `" value="` + value + `"/>
	<label class="form-check-label" for="` + id + `">` + text + `</label>
  </div>`)
	if len(e.children) == 0 {
		item.children[0].Attr("checked", "true")
	}
	e.add(item)

	return e
}

func (e *radioElement) Select(value string) *radioElement {
	// fmt.Println("set check:", value)
	e.Traverse(func(parent string, ht *HtmlToken) error {
		if ht.Info.Data != "input" {
			return nil
		}
		if ht.GetAttr("value") == value {
			// fmt.Println("success to set:", ht.String())
			ht.Attr("checked", "true")
		} else {
			ht.Attr("checked", "")
		}
		return nil
	})
	return e
}

// Inline Make items appear inline. Must be called after adding items.
func (e *radioElement) Inline() *radioElement {
	e.Traverse(func(parent string, ht *HtmlToken) error {
		if ht.Info.Data != "div" || parent != "div" {
			return nil
		}
		cls := ht.GetAttr("class")
		if !strings.Contains(cls, "form-check") {
			return nil
		}
		ht.Attr("class", "form-check form-check-inline")
		return nil
	})
	return e
}

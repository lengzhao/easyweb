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
	out.Parse("<div></div>")
	if name == "" {
		name = util.GetCallerID(util.LevelParent)
	}
	out.name = name
	out.Attr("id", getID())
	return &out
}

func (e *radioElement) AddItem(value, text string) *radioElement {
	id := getID()
	e.add(parseStringToToken(`<div class="form-check">
	<input class="form-check-input" type="radio" name="` + e.name + `" id="` + id + `" value="` + value + `"/>
	<label class="form-check-label" for="` + id + `">` + text + `</label>
  </div>`))

	return e
}

// Inline Make items appear inline. Must be called after adding items.
func (e *radioElement) Inline() *radioElement {
	e.Traverse(func(ht *HtmlToken) error {
		if ht.info.Data != "div" {
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

type RadioItem struct {
	Value string
	Text  string
}

func (e *radioElement) Add(in any) *radioElement {
	switch val := in.(type) {
	case map[string]string:
		for name, text := range val {
			e.AddItem(name, text)
		}
	case []RadioItem:
		for _, v := range val {
			e.AddItem(v.Value, v.Text)
		}
	case RadioItem:
		e.AddItem(val.Value, val.Text)
	default:
		e.add(Box(in))
	}
	return e
}

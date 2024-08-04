package e

import (
	"fmt"
	"sort"
)

type tableElement struct {
	HtmlToken
	boldFirst bool
}

var _ IElement = &tableElement{}

func Table(head ...string) *tableElement {
	var out tableElement
	out.parseText(`<table class="table table-bordered table-striped">
	<thead>
	  <tr>
	  </tr>
	</thead>
	<tbody>
	</tbody>
  </table>`)
	out.Traverse(nil, func(parent, ht IElement) error {
		if ht.HtmlTag() != "tr" {
			return nil
		}
		for _, v := range head {
			ht.Add(MustParseHtml(`<th scope="col">` + v + `</th>`))
		}
		return nil
	})
	out.SetAttr("id", getID())
	return &out
}

func (e *tableElement) AddLine(in []any) *tableElement {
	tr, _ := ParseHtml(`<tr></tr>`)
	for i, v := range in {
		td, _ := ParseHtml(`<td></td>`)
		if e.boldFirst && i == 0 {
			td, _ = ParseHtml(`<th scope="row"></th>`)
		}
		switch it := v.(type) {
		case IElement:
			td.Add(it)
		case string:
			td.SetText(it)
		default:
			str := fmt.Sprint(v)
			td.SetText(str)
		}
		tr.Add(td)
	}
	e.children[1].Add(tr)
	return e
}

func (e *tableElement) AddValue(in [][]any) *tableElement {
	for _, it := range in {
		e.AddLine(it)
	}

	return e
}

// func (e *tableElement) BoldFirstRow() *tableElement {
// 	e.boldFirst = true

// 	e.Traverse(nil, func(parent, ht IElement) error {
// 		if parent == nil || ht.HtmlTag() != "tbody" || parent.HtmlTag() != "table" {
// 			return nil
// 		}
// 		for _, it := range ht.GetChilds() {
// 			it.children[0].Info.Data = "th"
// 			it.children[0].SetAttr("scope", "row")
// 		}

// 		return fmt.Errorf("finish")
// 	})
// 	return e
// }

func (e *tableElement) HiddenHead(hidden bool) *tableElement {
	e.boldFirst = true

	e.Traverse(nil, func(parent, ht IElement) error {
		if parent == nil || ht.HtmlTag() != "thead" || parent.HtmlTag() != "table" {
			return nil
		}
		if hidden {
			ht.SetAttr("hidden", "true")
		} else {
			ht.SetAttr("hidden", "")
		}
		return fmt.Errorf("finish")
	})
	return e
}

// keyWidth: 0-12, total width is 12, 0:auto
func Map2Table(keyWidth int, in map[string]any) *tableElement {
	t := Table("Key", "Value")
	// t.BoldFirstRow()
	var keys []string
	for k := range in {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		item := in[k]
		switch val := item.(type) {
		case []any:
			element := List(val...).SetAttr("class", "list-group list-group-flush")
			t.AddLine([]any{k, element})
		case map[string]any:
			t.AddLine([]any{k, Map2Table(keyWidth, val).HiddenHead(true)})
		default:
			t.AddLine([]any{k, val})
		}
	}
	if keyWidth <= 0 || keyWidth >= 12 {
		return t
	}
	width := fmt.Sprintf("col-%d", keyWidth)
	t.Traverse(nil, func(parent, ht IElement) error {
		if ht.HtmlTag() != "tr" {
			return nil
		}
		child := ht.GetChilds()
		if len(child) == 0 {
			return nil
		}
		child[0].SetAttr("class", width)
		return nil
	})
	return t
}

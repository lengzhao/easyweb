package e

import (
	"fmt"
	"sort"
)

type tableElement struct {
	HtmlToken
	boldFirst bool
}

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
	out.Traverse(func(parent string, ht *HtmlToken) error {
		if ht.Info.Data != "tr" {
			return nil
		}
		for _, v := range head {
			ht.add(`<th scope="col">` + v + `</th>`)
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
		td.add(v)
		tr.add(td)
	}
	e.children[1].add(tr)
	return e
}

func (e *tableElement) AddValue(in [][]any) *tableElement {
	for _, it := range in {
		e.AddLine(it)
	}

	return e
}

func (e *tableElement) BoldFirstRow() *tableElement {
	e.boldFirst = true

	e.Traverse(func(parent string, ht *HtmlToken) error {
		if ht.Info.Data != "tbody" || parent != "table" {
			return nil
		}
		for _, it := range ht.children {
			it.children[0].Info.Data = "th"
			it.children[0].SetAttr("scope", "row")
		}

		return fmt.Errorf("finish")
	})
	return e
}

func (e *tableElement) HiddenHead() *tableElement {
	e.boldFirst = true

	e.Traverse(func(parent string, ht *HtmlToken) error {
		if ht.Info.Data != "thead" || parent != "table" {
			return nil
		}
		ht.children = nil
		return fmt.Errorf("finish")
	})
	return e
}

// keyWidth: 0-6, total width is 12, 0:auto,
func Map2Table(keyWidth int, in map[string]any) *tableElement {
	t := Table("Key", "Value")
	t.BoldFirstRow()
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
			t.AddLine([]any{k, Map2Table(keyWidth, val).HiddenHead()})
		default:
			t.AddLine([]any{k, in[k]})
		}
	}
	if keyWidth <= 0 || keyWidth >= 12 {
		return t
	}
	width := fmt.Sprintf("col-%d", keyWidth)
	t.Traverse(func(parent string, ht *HtmlToken) error {
		if ht.Info.Data != "tr" {
			return nil
		}
		if len(ht.children) == 0 {
			return nil
		}
		ht.children[0].SetAttr("class", width)
		return nil
	})
	return t
}

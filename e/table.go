package e

import "fmt"

type tableElement struct {
	HtmlToken
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
	out.Traverse(func(ht *HtmlToken) error {
		if ht.info.Data != "tr" {
			return nil
		}
		for _, v := range head {
			ht.add(`<th scope="col">` + v + `</th>`)
		}
		return nil
	})
	out.Attr("id", getID())
	return &out
}

func (e *tableElement) AddItem(in []any) *tableElement {
	tr, _ := ParseHtml(`<tr></tr>`)
	for _, v := range in {
		td, _ := ParseHtml(`<td></td>`)
		td.add(v)
		tr.add(td)
	}
	e.children[1].add(tr)
	return e
}

func (e *tableElement) AddValue(in [][]any) *tableElement {
	for _, it := range in {
		e.AddItem(it)
	}

	return e
}

func (e *tableElement) BoldFirstRow() *tableElement {

	e.Traverse(func(ht *HtmlToken) error {
		if ht.info.Data != "tbody" {
			return nil
		}
		for _, it := range ht.children {
			it.children[0].info.Data = "th"
			it.children[0].Attr("scope", "row")
		}

		return fmt.Errorf("finish")
	})
	return e
}

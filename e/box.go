package e

import "fmt"

func Box(items ...any) IElement {
	out := MustParseHtml(`<div></div>`)
	out.SetAttr("id", getID())
	for _, it := range items {
		switch v := it.(type) {
		case IElement:
			out.Add(v)
		case string:
			out.Add(MustParseHtml(`<div>` + v + "</div>"))
		default:
			str := fmt.Sprint(v)
			out.Add(MustParseHtml(`<div>` + str + "</div>"))
		}
	}
	return out
}

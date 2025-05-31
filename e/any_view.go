package e

import "fmt"

func AnyView(input any) IElement {
	out := MustParseHtml(`<div></div>`)
	out.SetAttr("id", getID())
	switch v := input.(type) {
	case IElement:
		out.Add(v)
	case []any:
		out.Add(Array2List(v...))
	case map[string]any:
		out.Add(Map2Table(0, v))
	case string:
		out.Add(MustParseHtml(`<div>` + v + "</div>"))
	default:
		str := fmt.Sprint(v)
		out.Add(MustParseHtml(`<div>` + str + "</div>"))
	}
	return out
}

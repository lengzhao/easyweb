package e

import "fmt"

// Col generates a col Element.
//
// Parameters:
//   - width: an integer representing the width of the colElement, 0-12, 0=auto.
//   - item: any item that needs to be added to the colElement.
func Col(width int, item IElement) IElement {
	cls := "col"
	if width > 0 && width <= 12 {
		cls = fmt.Sprint("col-", width)
	}
	out := MustParseHtml(`<div class="` + cls + `"></div>`)

	out.SetAttr("id", getID())
	out.Add(item)
	return out
}

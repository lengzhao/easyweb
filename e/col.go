package e

import "fmt"

type colElement struct {
	HtmlToken
}

// Col generates a col Element.
//
// Parameters:
//   - width: an integer representing the width of the colElement, 0-12, 0=auto.
//   - item: any item that needs to be added to the colElement.
func Col(width int, item any) *colElement {
	out := colElement{}
	cls := "col"
	if width > 0 && width <= 12 {
		cls = fmt.Sprint("col-", width)
	}

	out.parseText(`<div class="` + cls + `"></div>`)
	out.Attr("id", getID())
	out.add(item)
	return &out
}

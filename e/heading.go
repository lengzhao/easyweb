package e

import "fmt"

type headingElement struct {
	HtmlToken
}

// h1-h6,leven=1-6
func Heading(level uint, msg string) *headingElement {
	out := headingElement{}
	out.parseText(fmt.Sprintf("<h%d>%s</h%d>", level, msg, level))
	return &out
}

package e

import (
	"bytes"
	"html/template"
)

type HtmlNode struct {
	e          string
	attributes map[string]string
	text       string

	child []*HtmlNode
}

func NewNode(element string) *HtmlNode {
	return &HtmlNode{
		e:          element,
		attributes: make(map[string]string),
	}
}

func (n *HtmlNode) AddAttribute(k, v string) *HtmlNode {
	n.attributes[k] = v
	return n
}

func (n *HtmlNode) AddChild(child *HtmlNode) *HtmlNode {
	n.child = append(n.child, child)
	return n
}

func (n *HtmlNode) SetText(text string) *HtmlNode {
	buff := new(bytes.Buffer)
	template.HTMLEscape(buff, []byte(text))
	n.text = buff.String()
	return n
}

func (n *HtmlNode) SetHtml(text string) *HtmlNode {
	n.text = text
	return n
}

func (n *HtmlNode) String() string {
	prefix := "<" + n.e
	suffix := "</" + n.e + ">"
	for k, v := range n.attributes {
		if v != "" {
			prefix += " " + k + "=\"" + v + "\""
		} else {
			prefix += " " + k
		}
	}
	prefix += ">"
	prefix += n.text
	for _, c := range n.child {
		prefix += c.String()
	}
	prefix += suffix
	return prefix
}

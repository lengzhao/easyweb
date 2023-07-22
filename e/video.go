package e

import "github.com/lengzhao/easyweb/util"

type VideoElement struct {
	BaseElement
	title string
}

func Video(url string) *VideoElement {
	var out VideoElement
	out.cont = url
	out.id = util.GetID()
	return &out
}

func (b *VideoElement) SetTitle(in string) *VideoElement {
	b.title = in
	return b
}

// ratio-1x1,ratio-4x3,ratio-16x9,ratio-21x9
func (b *VideoElement) Class(in string) *VideoElement {
	b.cls += in
	return b
}

func (b *VideoElement) String() string {
	if b.cls == "" {
		b.cls = "ratio-16x9"
	}
	node := NewNode("div").AddAttribute("class", "ratio "+b.cls)
	if b.id != "" {
		node.AddAttribute("id", b.id)
	}
	node.AddChild(NewNode("iframe").AddAttribute("src", b.cont).AddAttribute("title", b.title).AddAttribute("allowfullscreen", ""))
	return node.String()
}

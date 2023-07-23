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
	node := NewNode("div").SetAttr("class", "ratio "+b.cls)
	if b.id != "" {
		node.SetAttr("id", b.id)
	}
	node.AddChild(NewNode("iframe").SetAttr("src", b.cont).SetAttr("title", b.title).SetAttr("allowfullscreen", ""))
	return node.String()
}

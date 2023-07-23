package e

import "github.com/lengzhao/easyweb/util"

type ImageElement struct {
	BaseElement
}

func Image(url string) *ImageElement {
	var out ImageElement
	out.cont = url
	out.id = util.GetID()
	return &out
}

// ratio-1x1,ratio-4x3,ratio-16x9,ratio-21x9
func (b *ImageElement) Class(in string) *ImageElement {
	b.cls += in
	return b
}

func (b *ImageElement) String() string {
	/*
		<div class="text-center">
		  <img src="..." class="rounded" alt="...">
		</div>
	*/
	if b.cls == "" {
		b.cls = "text-center"
	}
	node := NewNode("div").SetAttr("class", b.cls)
	if b.id != "" {
		node.SetAttr("id", b.id)
	}
	node.AddChild(NewNode("img").SetAttr("src", b.cont).SetAttr("class", "rounded").SetAttr("alt", b.cont))
	return node.String()
}

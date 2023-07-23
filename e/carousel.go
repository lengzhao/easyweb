package e

import (
	"fmt"

	"github.com/lengzhao/easyweb/util"
)

type CarouselElement struct {
	BaseElement
	slides []string
}

func Carousel(id string) *CarouselElement {
	var out CarouselElement
	if id == "" {
		id = util.GetCallerID(util.LevelParent)
	}
	out.id = id
	return &out
}

func (b *CarouselElement) SetImg(in []string) *CarouselElement {
	b.slides = in
	return b
}

func (b *CarouselElement) Class(in string) *CarouselElement {
	b.cls += in
	return b
}

func (b *CarouselElement) String() string {
	node := NewNode("div")
	node.SetAttr("class", "carousel slide "+b.cls)
	node.SetAttr("id", b.id)

	btns := NewNode("div").SetAttr("class", "carousel-indicators")
	for i := range b.slides {
		btn := NewNode("button").SetAttr("type", "button").SetAttr("data-bs-target", "#"+b.id).SetAttr("data-bs-slide-to", fmt.Sprint(i))
		if i == 0 {
			btn.SetAttr("class", "active")
		}
		btns.AddChild(btn)
	}
	node.AddChild(btns)
	imgs := NewNode("div").SetAttr("class", "carousel-inner")
	for i, s := range b.slides {
		img := NewNode("div")
		if i == 0 {
			img.SetAttr("class", "carousel-item active")
		} else {
			img.SetAttr("class", "carousel-item")
		}
		img.AddChild(NewNode("img").SetAttr("src", s).SetAttr("class", "d-block w-100"))
		imgs.AddChild(img)
	}
	node.AddChild(imgs)
	btn1 := NewNode("button").SetAttr("class", "carousel-control-prev").SetAttr("type", "button").SetAttr("data-bs-target", "#"+b.id).SetAttr("data-bs-slide", "prev")
	btn1.AddChild(NewNode("span").SetAttr("class", "carousel-control-prev-icon").SetAttr("aria-hidden", "true"))
	btn1.AddChild(NewNode("span").SetAttr("class", "visually-hidden").SetText("Previous"))
	node.AddChild(btn1)
	btn2 := NewNode("button").SetAttr("class", "carousel-control-next").SetAttr("type", "button").SetAttr("data-bs-target", "#"+b.id).SetAttr("data-bs-slide", "next")
	btn2.AddChild(NewNode("span").SetAttr("class", "carousel-control-next-icon").SetAttr("aria-hidden", "true"))
	btn2.AddChild(NewNode("span").SetAttr("class", "visually-hidden").SetText("Next"))
	node.AddChild(btn2)

	return node.String()
}

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
	node.AddAttribute("class", "carousel slide "+b.cls)
	node.AddAttribute("id", b.id)

	btns := NewNode("div").AddAttribute("class", "carousel-indicators")
	for i := range b.slides {
		btn := NewNode("button").AddAttribute("type", "button").AddAttribute("data-bs-target", "#"+b.id).AddAttribute("data-bs-slide-to", fmt.Sprint(i))
		if i == 0 {
			btn.AddAttribute("class", "active")
		}
		btns.AddChild(btn)
	}
	node.AddChild(btns)
	imgs := NewNode("div").AddAttribute("class", "carousel-inner")
	for i, s := range b.slides {
		img := NewNode("div")
		if i == 0 {
			img.AddAttribute("class", "carousel-item active")
		} else {
			img.AddAttribute("class", "carousel-item")
		}
		img.AddChild(NewNode("img").AddAttribute("src", s).AddAttribute("class", "d-block w-100"))
		imgs.AddChild(img)
	}
	node.AddChild(imgs)
	btn1 := NewNode("button").AddAttribute("class", "carousel-control-prev").AddAttribute("type", "button").AddAttribute("data-bs-target", "#"+b.id).AddAttribute("data-bs-slide", "prev")
	btn1.AddChild(NewNode("span").AddAttribute("class", "carousel-control-prev-icon").AddAttribute("aria-hidden", "true"))
	btn1.AddChild(NewNode("span").AddAttribute("class", "visually-hidden").SetText("Previous"))
	node.AddChild(btn1)
	btn2 := NewNode("button").AddAttribute("class", "carousel-control-next").AddAttribute("type", "button").AddAttribute("data-bs-target", "#"+b.id).AddAttribute("data-bs-slide", "next")
	btn2.AddChild(NewNode("span").AddAttribute("class", "carousel-control-next-icon").AddAttribute("aria-hidden", "true"))
	btn2.AddChild(NewNode("span").AddAttribute("class", "visually-hidden").SetText("Next"))
	node.AddChild(btn2)

	return node.String()
}

package e

import (
	"fmt"
)

type CardElement struct {
	BaseElement
	img      string
	link     string
	linkText string
	title    string
	subTitle string
}

func Card(text string) *CardElement {
	var out CardElement
	out.cont = text
	return &out
}

func (b *CardElement) Image(in string) *CardElement {
	b.img = in
	return b
}

func (b *CardElement) Link(url, text string) *CardElement {
	b.link = url
	b.linkText = text
	return b
}

func (b *CardElement) Text(in any) *CardElement {
	b.cont = fmt.Sprint(in)
	return b
}

func (b *CardElement) Title(title, subTitle string) *CardElement {
	b.title = title
	b.subTitle = subTitle
	return b
}

func (b *CardElement) String() string {
	node := NewNode("div")
	node.AddAttribute("class", "card")
	if b.id != "" {
		node.AddAttribute("id", b.id)
	}
	if b.img != "" {
		img := NewNode("img").AddAttribute("src", b.img).AddAttribute("class", "card-img-top")
		node.AddChild(img)
	}
	body := NewNode("div").AddAttribute("class", "card-body")
	if b.title != "" {
		title := NewNode("h5").AddAttribute("class", "card-title").SetText(b.title)
		body.AddChild(title)
	}
	if b.subTitle != "" {
		subTitle := NewNode("p").AddAttribute("class", "card-text").SetText(b.subTitle)
		body.AddChild(subTitle)
	}
	if b.cont != "" {
		cont := NewNode("p").AddAttribute("class", "card-text").SetText(b.cont)
		body.AddChild(cont)
	}
	if b.link != "" {
		link := NewNode("a").AddAttribute("href", b.link).AddAttribute("class", "btn btn-primary").SetText(b.linkText)
		body.AddChild(link)
	}
	node.AddChild(body)
	return node.String()
}

package e

import (
	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/util"
)

type ButtonElement struct {
	BaseElement
	cb func(id string)
}

var _ easyweb.Event = &ButtonElement{}

func Button(text string, cb func(id string)) *ButtonElement {
	out := &ButtonElement{}
	out.cont = text
	out.cb = cb
	out.id = util.GetID()
	return out
}

// MessageCb only call by framework
func (b *ButtonElement) MessageCb(id, info string) {
	if b.cb != nil {
		b.cb(id)
	}
}

func (b *ButtonElement) String() string {
	node := NewNode("button")
	if b.cls == "" {
		b.cls = "btn-primary"
	}
	node.SetAttr("class", "btn "+b.cls)
	if b.id != "" {
		node.SetAttr("id", b.id)
	}
	node.SetText(b.cont)
	return node.String()
}

// btn-primary
func (b *ButtonElement) Class(in string) *ButtonElement {
	b.cls += in
	return b
}

func (b *ButtonElement) EventInfo() (string, string) {
	return b.id, string(EventButton)
}

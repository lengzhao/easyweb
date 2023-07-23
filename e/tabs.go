package e

import (
	"fmt"

	"github.com/lengzhao/easyweb/util"
)

/*
<nav>
  <div class="nav nav-tabs" id="nav-tab" role="tablist">
    <button class="nav-link active" id="nav-home-tab" data-bs-toggle="tab" data-bs-target="#nav-home" type="button" role="tab" aria-controls="nav-home" aria-selected="true">Home</button>
    <button class="nav-link" id="nav-profile-tab" data-bs-toggle="tab" data-bs-target="#nav-profile" type="button" role="tab" aria-controls="nav-profile" aria-selected="false">Profile</button>
    <button class="nav-link" id="nav-contact-tab" data-bs-toggle="tab" data-bs-target="#nav-contact" type="button" role="tab" aria-controls="nav-contact" aria-selected="false">Contact</button>
  </div>
</nav>
<div class="tab-content" id="nav-tabContent">
  <div class="tab-pane fade show active" id="nav-home" role="tabpanel" aria-labelledby="nav-home-tab">...</div>
  <div class="tab-pane fade" id="nav-profile" role="tabpanel" aria-labelledby="nav-profile-tab">...</div>
  <div class="tab-pane fade" id="nav-contact" role="tabpanel" aria-labelledby="nav-contact-tab">...</div>
</div>
*/

type TabsElement struct {
	// BaseElement
	node *HtmlNode
}

func Tabs() *TabsElement {
	var out TabsElement
	out.node = NewNode("div")
	out.node.SetAttr("id", util.GetID())
	head := NewNode("nav").SetAttr("role", "tablist").SetAttr("class", "nav nav-tabs")
	out.node.AddChild(NewNode("nav").AddChild(head))
	body := NewNode("div").SetAttr("class", "tab-content").SetAttr("id", util.GetID())
	out.node.AddChild(body)

	return &out
}

func (b *TabsElement) Add(title string, body any) *TabsElement {
	head := b.node.GetChild(0, 0)
	lb := b.node.GetChild(1)
	first := head.GetChild(0) == nil
	//<button class="nav-link active" id="nav-home-tab" data-bs-toggle="tab" data-bs-target="#nav-home" type="button" role="tab" aria-controls="nav-home" aria-selected="true">Home</button>
	id := util.GetID()
	bid := util.GetID()
	hItem := NewNode("button").SetAttr("id", id).SetAttr("type", "button").SetAttr("role", "tab").SetAttr("aria-controls", bid).SetAttr("data-bs-toggle", "tab")
	hItem.SetAttr("data-bs-target", "#"+bid)
	if first {
		hItem.SetAttr("class", "nav-link active").SetAttr("aria-selected", "true")
	} else {
		hItem.SetAttr("class", "nav-link").SetAttr("aria-selected", "false")
	}
	hItem.SetText(title)

	//<div class="tab-pane fade show active" id="nav-home" role="tabpanel" aria-labelledby="nav-home-tab">...</div>
	bItem := NewNode("div").SetAttr("id", bid).SetAttr("role", "tabpanel").SetAttr("aria-labelledby", id)
	if first {
		bItem.SetAttr("class", "tab-pane fade show active")
	} else {
		bItem.SetAttr("class", "tab-pane fade")
	}
	bItem.SetHtml(fmt.Sprint(body))
	head.AddChild(hItem)
	lb.AddChild(bItem)

	return b
}

func (b *TabsElement) Class(in string) *TabsElement {
	cls := b.node.GetAttr("class")
	if cls == "" {
		cls = in
	} else {
		cls += " " + in
	}
	b.node.SetAttr("class", cls)
	return b
}

func (b *TabsElement) String() string {
	return b.node.String()
}

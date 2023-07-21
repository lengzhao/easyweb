package e

type ButtonElement struct {
	BaseElement
	cb func(id string)
}

// var _ Element = &ButtonElement{}

func Button(text string) *ButtonElement {
	out := &ButtonElement{}
	out.cont = text

	return out
}

// ElementCb only call by framework
func (b *ButtonElement) ElementCb(id, info string) {
	// fmt.Println("button event:", id, info)
	if b.cb != nil {
		b.cb(id)
	}
}

func (b *ButtonElement) String() string {
	node := NewNode("button")
	if b.cls == "" {
		b.cls = "btn-primary"
	}
	node.AddAttribute("class", "btn "+b.cls)
	node.SetText(b.cont)
	return node.String()
}

func (b *ButtonElement) SetCb(cb func(id string)) *ButtonElement {
	b.cb = cb
	return b
}

// btn-primary
func (b *ButtonElement) Class(in string) *ButtonElement {
	b.cls += in
	return b
}

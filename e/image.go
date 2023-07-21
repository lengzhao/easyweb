package e

type ImageElement struct {
	BaseElement
}

func Image(url string) *ImageElement {
	var out ImageElement
	out.cont = url
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
	node := NewNode("div").AddAttribute("class", b.cls)
	node.AddChild(NewNode("img").AddAttribute("src", b.cont).AddAttribute("class", "rounded").AddAttribute("alt", b.cont))
	return node.String()
}

package e

type VideoElement struct {
	BaseElement
	title string
}

func Video(url string) *VideoElement {
	var out VideoElement
	out.cont = url
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
	/*
		<div class="ratio ratio-16x9">
		  <iframe src="https://www.youtube.com/embed/zpOULjyy-n8?rel=0" title="YouTube video" allowfullscreen></iframe>
		</div>
	*/
	if b.cls == "" {
		b.cls = "ratio-16x9"
	}
	node := NewNode("div").AddAttribute("class", "ratio "+b.cls)
	node.AddChild(NewNode("iframe").AddAttribute("src", b.cont).AddAttribute("title", b.title).AddAttribute("allowfullscreen", ""))
	return node.String()
}

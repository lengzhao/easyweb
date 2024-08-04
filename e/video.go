package e

type videoElement struct {
	HtmlToken
}

var _ IElement = &videoElement{}

// Video generates a video HTML element from the given URL. defults to Aspect16by9
func Video(url string) *videoElement {
	var out videoElement
	out.parseText(`<div class="embed-responsive embed-responsive-16by9">
	<iframe class="embed-responsive-item" src="` + url + `" allowfullscreen></iframe>
  </div>`)
	out.SetAttr("id", getID())
	return &out
}

func (e *videoElement) Aspect21by9() *videoElement {
	e.SetAttr("class", "embed-responsive embed-responsive-21by9")
	return e
}
func (e *videoElement) Aspect4by3() *videoElement {
	e.SetAttr("class", "embed-responsive embed-responsive-4by3")
	return e
}

func (e *videoElement) Aspect1by1() *videoElement {
	e.SetAttr("class", "embed-responsive embed-responsive-1by1")
	return e
}

func (e *videoElement) Aspect16by9() *videoElement {
	e.SetAttr("class", "embed-responsive embed-responsive-16by9")
	return e
}

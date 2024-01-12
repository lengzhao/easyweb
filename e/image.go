package e

type imageElement struct {
	HtmlToken
}

func Image(url string) *imageElement {
	var out imageElement
	out.parseText(`<img src="` + url + `" class="img-fluid" alt="...">`)
	out.SetAttr("id", getID())
	return &out
}

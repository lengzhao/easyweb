package e

func Image(url string) IElement {
	out := MustParseHtml(`<img src="` + url + `" class="img-fluid" alt="...">`)
	out.SetAttr("id", getID())
	return out
}

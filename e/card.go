package e

import (
	"fmt"
)

type cardElement struct {
	HtmlToken
}

func (c cardElement) Data() string {
	return `<div class="card col-3">
	<img src="..." class="card-img-top" alt="..."/>
	<div class="card-body">
	  <h5 class="card-title"></h5>
	  <h6 class="card-subtitle mb-2 text-muted"></h6>
	  <p class="card-text"></p>
	  <a href="#" class="btn btn-primary">Go somewhere</a>
	</div>
  </div>`
}

// 通过调用初始化的时候，添加参数，参数类型为ParamItem，包含key和value

func Card() *cardElement {
	out := cardElement{}
	out.parseText(out.Data())
	out.Traverse(nil, func(parent IElement, ht IElement) error {
		switch ht.HtmlToken().Data {
		case "img", "a", "p", "h5", "h6":
			ht.SetAttr("hidden", "true")
		}
		return nil
	})
	out.SetAttr("id", getID())
	return &out
}

func (b *cardElement) Image(src, alt string) *cardElement {
	if src == "" {
		return b
	}
	b.Traverse(nil, func(parent, ht IElement) error {
		if ht.HtmlToken().Data == "img" {
			ht.SetAttr("src", src)
			ht.SetAttr("alt", alt)
			ht.SetAttr("hidden", "")
			return fmt.Errorf("finish")
		}
		return nil
	})
	return b
}

func (b *cardElement) Title(title, subTitle string) *cardElement {
	b.Traverse(nil, func(parent, ht IElement) error {
		if ht.HtmlToken().Data == "h5" {
			ht.SetChild()
			ht.SetText(title)
			ht.SetAttr("hidden", "")
		}
		if ht.HtmlToken().Data == "h6" {
			ht.SetChild()
			ht.SetText(title)
			ht.SetAttr("hidden", "")
			return fmt.Errorf("finish")
		}
		return nil
	})
	return b
}

func (b *cardElement) Link(url, text string) *cardElement {
	b.Traverse(nil, func(parent, ht IElement) error {
		if ht.HtmlToken().Data == "a" {
			ht.SetAttr("href", url)
			ht.SetChild()
			ht.SetText(text)
			ht.SetAttr("hidden", "")
			return fmt.Errorf("finish")
		}
		return nil
	})
	return b
}

func (b *cardElement) Text(in any) *cardElement {
	b.Traverse(nil, func(parent, ht IElement) error {
		if ht.HtmlToken().Data == "p" {
			ht.SetChild()
			ht.SetText(fmt.Sprint(in))
			ht.SetAttr("hidden", "")
			return fmt.Errorf("finish")
		}
		return nil
	})
	return b
}

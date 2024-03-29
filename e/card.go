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
	out.Traverse(func(parent string, ht *HtmlToken) error {
		switch ht.Info.Data {
		case "img", "a", "p", "h5", "h6":
			ht.disable = true
		}
		return nil
	})
	out.SetAttr("id", getID())
	out.disable = false
	return &out
}

func (b *cardElement) Image(src, alt string) *cardElement {
	if src == "" {
		return b
	}
	b.Traverse(func(parent string, ht *HtmlToken) error {
		if ht.Info.Data == "img" {
			ht.SetAttr("src", src)
			ht.SetAttr("alt", alt)
			ht.disable = false
			return fmt.Errorf("finish")
		}
		return nil
	})
	return b
}

func (b *cardElement) Title(title, subTitle string) *cardElement {
	b.Traverse(func(parent string, ht *HtmlToken) error {
		if ht.Info.Data == "h5" {
			ht.text = title
			ht.children = nil
			ht.disable = false
		}
		if ht.Info.Data == "h6" {
			ht.text = subTitle
			ht.children = nil
			ht.disable = false
			return fmt.Errorf("finish")
		}
		return nil
	})
	return b
}

func (b *cardElement) Link(url, text string) *cardElement {
	b.Traverse(func(parent string, ht *HtmlToken) error {
		if ht.Info.Data == "a" {
			ht.SetAttr("href", url)
			ht.text = text
			ht.children = nil
			ht.disable = false
			return fmt.Errorf("finish")
		}
		return nil
	})
	return b
}

func (b *cardElement) Text(in any) *cardElement {
	b.Traverse(func(parent string, ht *HtmlToken) error {
		if ht.Info.Data == "p" {
			ht.text = fmt.Sprint(in)
			ht.children = nil
			ht.disable = false
			return fmt.Errorf("finish")
		}
		return nil
	})
	return b
}

package e

import (
	"fmt"
)

type cardElement struct {
	HtmlToken
}

func (c cardElement) Data() string {
	return `<div class="card" style="width: 18rem;">
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

func Card1() *cardElement {
	out := cardElement{}
	out.Parse(out.Data())
	out.Traverse(func(ht *HtmlToken) error {
		switch ht.info.Data {
		case "img", "a", "p", "h5", "h6":
			ht.disable = true
		}
		return nil
	})
	out.Attr("id", getID())
	out.disable = false
	return &out
}

func (b *cardElement) Image(src, alt string) *cardElement {
	if src == "" {
		return b
	}
	b.Traverse(func(ht *HtmlToken) error {
		if ht.info.Data == "img" {
			ht.Attr("src", src)
			ht.Attr("alt", alt)
			ht.disable = false
			return fmt.Errorf("finish")
		}
		return nil
	})
	return b
}

func (b *cardElement) Title(title, subTitle string) *cardElement {
	b.Traverse(func(ht *HtmlToken) error {
		if ht.info.Data == "h5" {
			ht.text = title
			ht.children = nil
			ht.disable = false
		}
		if ht.info.Data == "h6" {
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
	b.Traverse(func(ht *HtmlToken) error {
		if ht.info.Data == "a" {
			ht.Attr("href", url)
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
	b.Traverse(func(ht *HtmlToken) error {
		if ht.info.Data == "p" {
			ht.text = fmt.Sprint(in)
			ht.children = nil
			ht.disable = false
			return fmt.Errorf("finish")
		}
		return nil
	})
	return b
}

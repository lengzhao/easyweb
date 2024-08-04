package e

import (
	"fmt"
)

type carouselElement struct {
	HtmlToken
	id string
}

func Carousel() *carouselElement {
	var out carouselElement
	out.id = getID()
	out.parseText(`<div id="` + out.id + `" class="carousel slide" data-bs-ride="carousel">
    <div class="carousel-indicators"></div>
    <div class="carousel-inner">
	</div>
    <button class="carousel-control-prev" type="button" data-bs-target="#` + out.id + `" data-bs-slide="prev">
        <span class="carousel-control-prev-icon" aria-hidden="true"></span>
        <span class="visually-hidden">Previous</span>
    </button>
    <button class="carousel-control-next" type="button" data-bs-target="#` + out.id + `" data-bs-slide="next">
        <span class="carousel-control-next-icon" aria-hidden="true"></span>
        <span class="visually-hidden">Next</span>
    </button>
</div>`)

	return &out
}

func (b *carouselElement) Add(image, title, body string) *carouselElement {
	indicators := b.children[0]
	ic := indicators.GetChilds()
	hd, _ := ParseHtml(`<button type="button" data-bs-target="#" data-bs-slide-to="0" aria-label="Slide ` + fmt.Sprintf("%d", len(ic)+1) + `"></button>`)
	if len(ic) == 0 {
		hd.SetAttr("class", "active")
		hd.SetAttr("aria-current", "true")
	}
	hd.SetAttr("data-bs-target", "#"+b.id)
	hd.SetAttr("data-bs-slide-to", fmt.Sprintf("%d", len(ic)))
	hd.SetAttr("aria-label", fmt.Sprintf("Slide %d", len(ic)+1))
	indicators.Add(hd)

	inner := b.children[1]
	bd, _ := ParseHtml(`<div class="carousel-item">
            <img src="` + image + `" class="d-block w-100" alt="..."/>
            <div class="carousel-caption d-none d-md-block">
                <h5>` + title + `</h5>
                <p>` + body + `</p>
            </div>
        </div>`)
	if len(ic) == 0 {
		bd.SetAttr("class", "carousel-item active")
	}
	bd.Traverse(nil, func(parent, ht IElement) error {
		if ht.HtmlTag() == "h5" {
			if title == "" {
				ht.SetAttr("hidden", "true")
			}
		}
		if ht.HtmlTag() == "p" {
			if body == "" {
				ht.SetAttr("hidden", "true")
			}
		}
		return nil
	})
	inner.Add(bd)
	return b
}

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

	hd, _ := ParseHtml(`<button type="button" data-bs-target="#" data-bs-slide-to="0" aria-label="Slide ` + fmt.Sprintf("%d", len(indicators.children)+1) + `"></button>`)
	if len(indicators.children) == 0 {
		hd.Attr("class", "active")
		hd.Attr("aria-current", "true")
	}
	hd.Attr("data-bs-target", "#"+b.id)
	hd.Attr("data-bs-slide-to", fmt.Sprintf("%d", len(indicators.children)))
	hd.Attr("aria-label", fmt.Sprintf("Slide %d", len(indicators.children)+1))
	indicators.add(hd)

	inner := b.children[1]
	bd, _ := ParseHtml(`<div class="carousel-item">
            <img src="` + image + `" class="d-block w-100" alt="..."/>
            <div class="carousel-caption d-none d-md-block">
                <h5>` + title + `</h5>
                <p>` + body + `</p>
            </div>
        </div>`)
	if len(inner.children) == 0 {
		bd.Attr("class", "carousel-item active")
	}
	bd.Traverse(func(ht *HtmlToken) error {
		if ht.info.Data == "h5" {
			if title == "" {
				ht.disable = true
			}
		}
		if ht.info.Data == "p" {
			if body == "" {
				ht.disable = true
			}
		}
		return nil
	})
	inner.add(bd)
	return b
}

package e

func Loading() IElement {
	item, _ := ParseHtml(`<div class="spinner-border text-info" role="status">
	<span class="visually-hidden">Loading...</span>
  </div>`)
	item.SetAttr("id", getID())
	return item
}

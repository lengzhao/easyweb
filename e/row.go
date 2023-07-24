package e

func Row(childs ...any) *boxElement {
	box := Box(childs...)
	box.Attr("class", "row")
	return box
}

package e

import (
	"testing"
)

func TestBox(t *testing.T) {
	it := Box("a", "b", "c")
	it.SetAttr("id", "")
	if it.String() != `<div>abc</div>` {
		t.Error("0", it.String())
	}
	it.Add(Div("1"), Div("2"), Div("3"))
	if it.String() != `<div>abc123</div>` {
		t.Error("1", it.String())
	}
	t.Log(it.String())
	// t.Error("aaa")
}

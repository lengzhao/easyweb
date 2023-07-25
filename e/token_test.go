package e

import (
	"testing"
)

func TestHtmlToken_Parse(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		// TODO: Add test cases.
		{"1", "<div></div>", "<div></div>"},
		{"2", "<div>a</div>", "<div>a</div>"},
		{"3", `<button/>`, `<button/>`},
		{"4", `<br>`, `<br></br>`},
		{"5", `<p>ssss</p>`, `<p>ssss</p>`},
		{"6", `<div><p>ssss</p><p>ssss</p></div>`, `<div><p>ssss</p><p>ssss</p></div>`},
		{"7", `<div id="">abc</div>`, `<div>abc</div>`},
		// input checkbox checked
		{"8", `<input type="checkbox" id="" checked/>`, `<input type="checkbox" checked/>`},
		// only text
		{"9", `hello1`, `hello1`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &HtmlToken{}
			n.parseText(tt.text)
			got := n.String()
			if got != tt.want {
				t.Errorf("name:%s, Parse() = %s, want %s", tt.name, got, tt.want)
			}
		})
	}
}

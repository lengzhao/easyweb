package e

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"github.com/lengzhao/easyweb/util"
)

type AccordionElement struct {
	BaseElement
}

func Accordion(id string) *AccordionElement {
	var out AccordionElement
	if id != "" {
		out.id = id
	} else {
		out.id = util.GetCallerID(util.LevelParent)
	}
	return &out
}

func (e *AccordionElement) String() string {
	node := NewNode("div")
	node.AddAttribute("class", "accordion "+e.cls)
	node.AddAttribute("id", e.id)
	node.SetHtml(e.cont)
	return node.String()
}

func (e *AccordionElement) Class(in string) *AccordionElement {
	e.cls += " " + in
	return e
}

type AccordionItem struct {
	Header string
	Text   string
}

func (e *AccordionElement) Add(in any) *AccordionElement {
	tempText := `<div class="accordion-item">
    <h2 class="accordion-header">
      <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#{{.ItemID}}" aria-expanded="true" aria-controls="{{.ItemID}}">
        {{.Header}}
      </button>
    </h2>
    <div id="{{.ItemID}}" class="accordion-collapse collapse show" data-bs-parent="#{{.Name}}">
      <div class="accordion-body">
        {{.Text}}
      </div>
    </div>
  </div>`
	type itemData struct {
		Name   string
		ItemID string
		Header string
		Text   string
	}
	switch val := in.(type) {
	case AccordionItem:
		data := itemData{
			Name:   e.id,
			ItemID: e.id + fmt.Sprint(e.cont),
			Header: val.Header,
			Text:   val.Text,
		}
		buff := new(bytes.Buffer)
		t, err := template.New("1").Parse(tempText)
		if err != nil {
			log.Println("fail to parse template:", err)
			return e
		}
		t.Execute(buff, data)
		e.cont += buff.String()
	case []AccordionItem:
		for _, v := range val {
			data := itemData{
				Name:   e.id,
				ItemID: e.id + fmt.Sprint(e.cont),
				Header: v.Header,
				Text:   v.Text,
			}
			buff := new(bytes.Buffer)
			t, err := template.New("1").Parse(tempText)
			if err != nil {
				log.Println("fail to parse template:", err)
				return e
			}
			t.Execute(buff, data)
			e.cont += buff.String()
		}
	default:
		e.cont += fmt.Sprint(in)
	}
	return e
}

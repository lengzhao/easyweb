package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/e"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	easyweb.RegisterPage(func(page easyweb.Page) {
		navbar := e.Navbar("Easy Web")
		navbar.Add("Home", "#")
		navbar.Add("Demo", "#")
		navbar.Add("Link", "#")
		navbar.Add("Action", "#")
		navbar.SetSearchCb(func(value string) {
			fmt.Println("search:", value)
		})
		page.Write(navbar)

		page.Title("Demo")
		tabs := e.Tabs()
		tabs.Add("Label", e.Box("<h1>Hello World</h1>", "<h2>Hello World</h2>", "<h3>Hello World</h3>"))
		tabs.Add("Accordion", e.Accordion().AddItem("H1", "<h1>Hello World</h1>").AddItem("H2", "<h2>Hello World</h2>").AddItem("H3", "<h3>Hello World</h3>").AddItem("H3-2", "<h3>Hello World</h3>"))
		tabs.Add("SVG", `<svg xmlns="http://www.w3.org/2000/svg" version="1.1">
							<circle cx="100" cy="50" r="40" stroke="black"
							stroke-width="2" fill="red" />
						</svg>`)

		box := e.Box()
		box.Add(e.InputGroup("input01", "input-group"))
		box.Add(e.InputGroup("name2", "Hello").Suffix("$$$$").Value("100000"))
		box.Add(e.InputGroup("name2", "Number").ChangeType("number"))
		box.Add(e.InputGroup("name2", "File").ChangeType("file"))
		box.Add(e.Radio("radio1").Add("name1", "text1").Add("name2", "text2").Add("name3", "text3").Inline().Select("name2"))
		list := e.List(e.Check("check1", "check text1").SetChecked(), e.Check("check2", "check text2"), e.Check("check3", "check text3")).Horizontal().ShowIndex()
		box.Add(list)
		box.Add(e.RangeInput("range", "range").SetRange(1, 30, 1).SetValue(15))

		tabs.Add("Input", box)

		var count int
		row := e.Box(e.Button("Click", func(id string) {
			fmt.Println("button click001:", id)
			page.Write("button click:" + fmt.Sprint(count))
			page.Write(time.Now().String())
		}))
		row.Add(e.Link("Link ...", "#"))
		dropdown := e.Dropdown("Dropdown001").AddLink("Link2 ...", "#")
		dropdown.AddButton("button3", func(id string) {
			page.Write("button3 click:" + id + time.Now().String())
		})
		dropdown.AddButton("button4", func(id string) {
			page.Write("button4 click:" + id + time.Now().String())
		})
		row.Add(dropdown)
		pag := e.Pagination([]string{"<<", "1", "2", "3", "4", "5", ">>"}, func(id, item string) {
			fmt.Println("pagination:", id, item)
		}).Active("2")
		row.Add(pag)
		tabs.Add("Button", row)

		box2 := e.Row()
		box2.Add(e.Card().Image("/static/1.png", "svg1").Title("Card1", "subtitle").Text("this is a card").Link("#", "Go ..."))
		box2.Add(e.Card().Image("/static/2.jpeg", "svg2").Title("Card2", "subtitle").Text("this is a card"))
		box2.Add(e.Card().Image("/static/3.jpeg", "svg3"))
		tabs.Add("Card", box2)

		carousel := e.Carousel()
		carousel.Add("/static/1.png", "image1", "this is a image:1")
		carousel.Add("/static/2.jpeg", "image2", "this is a image:2")
		carousel.Add("/static/3.jpeg", "image3", "this is a image:3")
		// carousel.Add("/static/kkk.jpg", "test", "lost the image")
		tabs.Add("Carousel", carousel)

		table := e.Table("#", "First", "Last", "Handle")
		table.AddValue([][]any{{"1", "Mark", "Otto", "@mdo"}, {"2", "Jacob", "Thornton", "@fat"}, {"3", "Larry the Bird", "...", "@twitter"}, {"4", "Mark", "Otto", "@mdo"}})
		table.BoldFirstRow()
		tabs.Add("Table", table)

		entry := e.InputGroup("name", "Name")
		radio := e.Radio("radio1").Add("name1", "text1").Add("name2", "text2").Add("name3", "text3").Inline().Select("name2")
		selectItem := e.Select("select1").Add("s1", "s1").Add("s2", "s2").Add("s3", "s3").Select("s2")
		datalist := e.Datalist("datalist", "Datalist")
		datalist.Add("value1", "value2", "San Francisco", "New York")
		form := e.Form(func(id string, info map[string]string) {
			fmt.Println("form data:", info)
			page.Write("form data:")
			page.Write(info)
		}).Add(entry).Add(radio).Add(selectItem).Add(datalist)
		tabs.Add("Form", form)

		box3 := e.Row()
		box3.Add(e.Button("Click", func(id string) {
			page.Write("button click:" + id + time.Now().String())
		})).Add(`<br/>`)
		label := e.Label("Click Label")
		label.SetCb("click", func(id string, dataType easyweb.CbDataType, data []byte) {
			page.Write("label click:" + id + time.Now().String())
		})
		box3.Add(label).Add(`<br/>`)
		box3.Add(e.Button("RunJs From Server", func(id string) {
			page.RunJs(`console.log("hello world. 123456")`)
			page.Write(`run js:console.log("hello world. 123456")` + time.Now().String())

		})).Add(`<br/>`)
		box3.Add(e.Button("Off Event(unbind self)", func(id string) {
			page.Write(`event off:` + time.Now().String())
			page.RegistEvent(id, "click", nil)
		})).Add(`<br/>`)
		box3.Add(e.Button("Color Change(setAttribute from server)", func(id string) {
			page.SetAttr(id, "class", "btn btn-success")
			page.Write(`Color Change to btn-success`)
		}))
		tabs.Add("Event", box3)

		modal := e.Modal("Open Modal", "Hello Modal").SetBody("Modal Body")
		tabs.Add("Modal", modal)

		page.Write(tabs)

	}, easyweb.DefaultPagePath...)
	http.ListenAndServe(":8182", nil)

}

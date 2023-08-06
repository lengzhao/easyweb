package e

import (
	"fmt"
	"log"
	"strings"
	"sync/atomic"

	"github.com/lengzhao/easyweb"
	"github.com/lengzhao/easyweb/util"
	"golang.org/x/net/html"
)

type ICallback func(id string, dataType easyweb.CbDataType, data []byte)

type HtmlToken struct {
	Info      html.Token
	children  []*HtmlToken
	parent    string
	text      string
	eventType string
	cb        ICallback
	disable   bool
}

func ParseHtml(text string) (*HtmlToken, error) {
	var out HtmlToken
	err := out.parseText(text)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (n *HtmlToken) parseText(text string) error {
	tkn := html.NewTokenizer(strings.NewReader(text))
	err := n.parse(tkn)
	if err != nil {
		return err
	}
	if len(n.children) == 0 {
		log.Println("fail to parse:", err)
		n.children = nil
	} else if n.Info.Data == "" && n.text == "" {
		*n = *n.children[0]
	}
	return nil
}

func (n *HtmlToken) parse(tkn *html.Tokenizer) error {
	for {
		tt := tkn.Next()
		if tt == html.ErrorToken {
			return nil
		}
		switch tt {
		case html.StartTagToken:
			child := &HtmlToken{}
			child.Info = tkn.Token()
			n.children = append(n.children, child)
			if selfClosingTagToken[child.Info.Data] {
				continue
			}
			child.parse(tkn)

		case html.SelfClosingTagToken:
			child := &HtmlToken{}
			child.Info = tkn.Token()
			n.children = append(n.children, child)
		case html.EndTagToken:
			lt := tkn.Token()
			if n.Info.Data != lt.Data {
				if selfClosingTagToken[lt.Data] {
					return nil
				}
				fmt.Println("warning end:", n.Info, lt.Data)
				return fmt.Errorf("end tag mismatch,hope:%s,get:%s", n.Info.Data, lt.Data)
			}
			return nil
		case html.TextToken:
			n.text += strings.TrimSpace(string(tkn.Text()))
		default:
			// fmt.Println("tt:", tt, tkn.Token(), tkn.Text())
		}
	}
}
func (n *HtmlToken) String() string {
	if n.disable {
		return ""
	}
	if n.Info.Data == "" {
		return n.text
	}
	out := "<" + n.Info.Data
	for _, it := range n.Info.Attr {
		if booleanAttributes[it.Key] {
			out += " " + it.Key
		} else if it.Val != "" {
			out += " " + it.Key
			out += `="` + it.Val + `"`
		}
	}
	if n.Info.Type == html.SelfClosingTagToken {
		return out + "/>"
	}
	out += ">"
	for _, child := range n.children {
		out += child.String()
	}
	out += n.text
	out += "</" + n.Info.Data + ">"
	return out
}

func (n *HtmlToken) GetID() string {
	return n.GetAttr("id")
}

// get Attribute
func (n *HtmlToken) GetAttr(k string) string {
	for _, it := range n.Info.Attr {
		if it.Key == k {
			if it.Val == "" && booleanAttributes[k] {
				return "true"
			}
			return it.Val
		}
	}
	return ""
}

// set Attribute
func (n *HtmlToken) Attr(k, v string) *HtmlToken {
	attr := []html.Attribute{}
	if v != "" {
		attr = append(attr, html.Attribute{Key: k, Val: v})
	}
	for _, it := range n.Info.Attr {
		if it.Key != k {
			attr = append(attr, it)
		}
	}
	n.Info.Attr = attr
	return n
}

// add children or text
func (n *HtmlToken) add(in ...any) *HtmlToken {
	for _, it := range in {
		switch val := it.(type) {
		case []iSelf:
			for _, it := range val {
				n.children = append(n.children, it.Self())
			}
		case iSelf:
			n.children = append(n.children, val.Self())
		default:
			item := HtmlToken{}
			item.text = fmt.Sprint(it)
			n.children = append(n.children, &item)
		}
	}
	return n
}

type iSelf interface {
	Self() *HtmlToken
}

var _ iSelf = &HtmlToken{}

func (n *HtmlToken) SetCb(typ string, cb ICallback) *HtmlToken {
	if typ == "" {
		typ = getEventType2(n.Info.Data)
	}
	n.eventType = typ
	n.cb = cb
	if n.GetAttr("id") == "" {
		n.Attr("id", util.GetID())
	}
	// fmt.Println("set event callback:", n.GetAttr("id"), n.Info.Data, n.cb)
	return n
}

// Self returns the HtmlToken itself. easy 'Set' its subclasses, or will lost callback event
func (n *HtmlToken) Self() *HtmlToken {
	return n
}

func (n *HtmlToken) AfterElementLoadedFromFramwork(p easyweb.Page) {
	// fmt.Println("try regist event:", n.GetAttr("id"), n.Info.Data, n.cb)
	if n.cb != nil && n.GetAttr("id") != "" {
		p.RegistEvent(n.GetAttr("id"), n.eventType, n)
	}
	for _, child := range n.children {
		child.AfterElementLoadedFromFramwork(p)
	}
}

func getEventType2(in string) string {
	switch in {
	case "form":
		return "form"
	case "input":
		return "input"
	case "textarea":
		return "input"
	case "button":
		return "button"
	default:
		return ""
	}
}

func (n *HtmlToken) MessageCallbackFromFramwork(id string, dataType easyweb.CbDataType, data []byte) bool {
	if id == n.GetAttr("id") {
		if n.cb != nil {
			n.cb(id, dataType, data)
		}
		return true
	}
	for _, child := range n.children {
		if child.MessageCallbackFromFramwork(id, dataType, data) {
			return true
		}
	}
	return false
}

type ITraverseCb func(parent string, token *HtmlToken) error

func (n *HtmlToken) Traverse(cb ITraverseCb) error {
	num := len(n.children)
	err := cb(n.parent, n)
	if err != nil {
		return err
	}
	for i, child := range n.children {
		child.parent = n.Info.Data
		if i >= num {
			// new children
			break
		}
		err = child.Traverse(cb)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *HtmlToken) AddChild(child *HtmlToken) *HtmlToken {
	n.children = append(n.children, child)
	return n
}

func (n *HtmlToken) GetChilds() []*HtmlToken {
	return n.children
}

func (n *HtmlToken) GetText() string {
	return n.text
}

var booleanAttributes map[string]bool = map[string]bool{
	"allowfullscreen": true,
	"async":           true,
	"autofocus":       true,
	"autoplay":        true,
	"checked":         true,
	"controls":        true,
	"default":         true,
	"defer":           true,
	"disabled":        true,
	"formnovalidate":  true,
	"inert":           true,
	"ismap":           true,
	"itemscope":       true,
	"loop":            true,
	"multiple":        true,
	"muted":           true,
	"nomodule":        true,
	"novalidate":      true,
	"open":            true,
	"playsinline":     true,
	"readonly":        true,
	"required":        true,
	"reversed":        true,
	"selected":        true,
}

var selfClosingTagToken map[string]bool = map[string]bool{
	"area":   true,
	"base":   true,
	"br":     true,
	"col":    true,
	"embed":  true,
	"hr":     true,
	"img":    true,
	"input":  true,
	"keygen": true,
	"link":   true,
	"meta":   true,
	"param":  true,
	"source": true,
	"track":  true,
	"wbr":    true,
}

var lastId int64

func getID() string {
	val := atomic.AddInt64(&lastId, 1)
	return fmt.Sprintf("eid%04d", val)
}

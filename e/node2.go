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

type ICallback func(id string, data []byte)

type HtmlToken struct {
	info     html.Token
	children []*HtmlToken
	text     string
	cb       ICallback
	disable  bool
}

func parseStringToToken(text string) *HtmlToken {
	var out HtmlToken
	out.Parse(text)
	return &out
}

func (n *HtmlToken) Parse(text string) {
	tkn := html.NewTokenizer(strings.NewReader(text))
	err := n.parse(tkn)
	if err != nil || len(n.children) == 0 {
		log.Println("fail to parse:", err)
		n.children = nil
	} else if n.info.Data == "" && n.text == "" {
		*n = *n.children[0]
	}
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
			child.info = tkn.Token()
			n.children = append(n.children, child)
			child.parse(tkn)
		case html.SelfClosingTagToken:
			child := &HtmlToken{}
			child.info = tkn.Token()
			n.children = append(n.children, child)
		case html.EndTagToken:
			lt := tkn.Token()
			if n.info.Data != lt.Data {
				fmt.Println("warning end:", n.info, lt.Data)
				return fmt.Errorf("end tag mismatch,hope:%s,get:%s", n.info.Data, lt.Data)
			}
			return nil
		case html.TextToken:
			n.text += strings.TrimSpace(string(tkn.Text()))
		default:
			fmt.Println("tt:", tt, tkn.Token(), tkn.Text())
		}
	}
}
func (n *HtmlToken) String() string {
	if n.disable {
		return ""
	}
	if n.info.Data == "" {
		return n.text
	}
	out := "<" + n.info.Data
	for _, it := range n.info.Attr {
		if booleanAttributes[it.Key] {
			out += " " + it.Key
		} else if it.Val != "" {
			out += " " + it.Key
			out += `="` + it.Val + `"`
		}
	}
	if n.info.Type == html.SelfClosingTagToken {
		return out + "/>"
	}
	out += ">"
	for _, child := range n.children {
		out += child.String()
	}
	out += n.text
	out += "</" + n.info.Data + ">"
	return out
}

// get Attribute
func (n *HtmlToken) GetAttr(k string) string {
	for _, it := range n.info.Attr {
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
	for i, it := range n.info.Attr {
		if it.Key == k {
			it.Val = v
			n.info.Attr[i] = it
			return n
		}
	}
	n.info.Attr = append(n.info.Attr, html.Attribute{Key: k, Val: v})
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
			fmt.Println("add:", it)
			n.text += fmt.Sprint(it)
		}
	}
	return n
}

type iSelf interface {
	Self() *HtmlToken
}

var _ iSelf = &HtmlToken{}

// Set sets the value(children/text) of the HtmlToken.
func (n *HtmlToken) Set(in any) *HtmlToken {
	switch val := in.(type) {
	case []iSelf:
		n.children = nil
		for _, it := range val {
			n.children = append(n.children, it.Self())
		}
	case iSelf:
		n.children = []*HtmlToken{val.Self()}
	default:
		n.text = fmt.Sprint(in)
	}

	return n
}

func (n *HtmlToken) SetCb(cb ICallback) *HtmlToken {
	n.cb = cb
	if n.GetAttr("id") == "" {
		n.Attr("id", util.GetID())
	}
	fmt.Println("set event callback:", n.GetAttr("id"), n.info.Data, n.cb)
	return n
}

// Self returns the HtmlToken itself. easy 'Set' its subclasses, or will lost callback event
func (n *HtmlToken) Self() *HtmlToken {
	return n
}

func (n *HtmlToken) RegistEvent(p easyweb.Page) {
	fmt.Println("try regist event:", n.GetAttr("id"), n.info.Data, n.cb)
	if n.cb != nil && n.GetAttr("id") != "" {
		p.RegistEvent(n.GetAttr("id"), getEventType2(n.info.Data), n)
	}
	for _, child := range n.children {
		child.RegistEvent(p)
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
	// case "radio":
	// 	return "change"
	case "button":
		return "button"
	default:
		return ""
	}
}

func (n *HtmlToken) MessageCallbackFromFramwork(id string, data []byte) bool {
	if id == n.GetAttr("id") {
		if n.cb != nil {
			n.cb(id, data)
		}
		return true
	}
	for _, child := range n.children {
		if child.MessageCallbackFromFramwork(id, data) {
			return true
		}
	}
	return false
}

type ITraverseCb func(*HtmlToken) error

func (n *HtmlToken) Traverse(cb ITraverseCb) error {
	err := cb(n)
	if err != nil {
		return err
	}
	for _, child := range n.children {
		err = child.Traverse(cb)
		if err != nil {
			return err
		}
	}
	return nil
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

var lastId int64

func getID() string {
	val := atomic.AddInt64(&lastId, 1)
	return fmt.Sprintf("eid%04d", val)
}

package e

import (
	"encoding/json"

	"github.com/lengzhao/easyweb"
)

type BaseEvent struct {
	id  string
	typ string
	cb  func(id, info string)
}
type EventType string

const (
	EventInput  EventType = "input"
	EventButton EventType = "button"
	EventForm   EventType = "form"
)

var _ easyweb.Event = &BaseEvent{}

func ButtonEvent(id string, cb func(id string)) *BaseEvent {
	return CustomEvent(id, string(EventButton), func(id, info string) {
		cb(id)
	})
}

func FormEvent(id string, cb func(id string, info map[string]string)) *BaseEvent {
	return CustomEvent(id, string(EventForm), func(id, info string) {
		var info2 map[string]string
		json.Unmarshal([]byte(info), &info2)
		cb(id, info2)
	})
}

func CustomEvent(id, typ string, cb func(id, info string)) *BaseEvent {
	out := &BaseEvent{}
	out.id = id
	out.typ = typ
	out.cb = cb
	return out
}

func (e BaseEvent) EventInfo() (id, typ string) {
	return e.id, e.typ
}
func (e BaseEvent) MessageCb(id, info string) {
	if e.cb != nil {
		e.cb(id, info)
	}
}

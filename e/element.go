package e

type Element interface {
	Class(string) Element
	Write(any) Element
	String() string
	GetID() string
	GetType() string
	ElementCb(id, info string)
}

type BaseElement struct {
	cont string
	cls  string
	id   string
}

func (e *BaseElement) SetID(id string) {
	e.id = id
}

func (e BaseElement) GetID() string {
	return e.id
}

func (e BaseElement) GetType() string {
	return ""
}

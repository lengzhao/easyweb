package e

type Element interface {
	// Class(string) Element
	// Write(any) Element
	String() string
	GetID() string
}

type BaseElement struct {
	cont string
	cls  string
	id   string
}

func (e *BaseElement) String() string {
	return ""
}

func (e BaseElement) GetID() string {
	return e.id
}

func (e *BaseElement) SetID(id string) {
	e.id = id
}

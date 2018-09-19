package util

// Element ...
type Element interface{}

// Collection ...
type Collection interface {
	Size() int
	IsEmpty() bool
	Add(Element) bool
	Remove(Element) error
	Contains(Element) bool
	Clear()
}

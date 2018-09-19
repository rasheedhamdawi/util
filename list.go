package util

// List ...
type List interface {
	Collection
	Get(int) Element
	IndexOf(Element) int
	Set(int, Element) error
	AddAt(int, Element) error
}

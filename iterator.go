package util

// Iterator ...
type Iterator interface {
	Next() Element
	HasNext() bool
}

// Iterable ...
type Iterable interface {
	ForEach(func(int, Element))
	Iterator() Iterator
}

package util

type elements map[int]Element

type arrayListIterator struct {
	current int
	elements
}

func (a *arrayListIterator) Next() Element {
	defer func() {
		a.current++
	}()

	return a.elements[a.current]
}

func (a *arrayListIterator) HasNext() bool {
	if a.current >= len(a.elements) {
		return false
	}

	return true
}

// ArrayList ...
type ArrayList struct {
	arrayListIterator
	elements
}

// Size ...
func (a *ArrayList) Size() int {
	return len(a.elements)
}

// IsEmpty ...
func (a *ArrayList) IsEmpty() bool {
	if a.Size() > 0 {
		return false
	}

	return true
}

// Add ...
func (a *ArrayList) Add(e Element) bool {
	// check if the list is empty make new one
	if a.elements == nil {
		a.elements = make(elements)
	}

	// add a new element to the end of the list
	if a.elements[a.Size()] = e; a.elements[a.Size()-1] == e {
		return true
	}

	return false
}

// Contains ...
func (a *ArrayList) Contains(e Element) bool {
	for _, item := range a.elements {
		if item == e {
			return true
		}
	}

	return false
}

// Get ...
func (a *ArrayList) Get(i int) Element {
	return a.elements[i]
}

// IndexOf ...
func (a *ArrayList) IndexOf(e Element) int {

	for index, item := range a.elements {
		if item == e {
			return index
		}
	}
	return -1
}

// Clear ...
func (a *ArrayList) Clear() {
	a.elements = make(elements)
}

// Remove ...
func (a *ArrayList) Remove(e Element) error {
	delete(a.elements, a.IndexOf(e))

	return nil
}

// Set ...
func (a *ArrayList) Set(index int, e Element) error {
	a.elements[index] = e

	return nil
}

// AddAt ...
func (a *ArrayList) AddAt(index int, e Element) error {

	return nil
}

// ToArray ...
func (a *ArrayList) ToArray() []Element {
	elements := make([]Element, a.Size())

	for i, v := range a.elements {
		elements[i] = v
	}

	return elements
}

// ForEach ...
func (a *ArrayList) ForEach(cb func(int, Element)) {
	for i, v := range a.elements {
		cb(i, v)
	}
}

// Iterator ...
func (a *ArrayList) Iterator() Iterator {
	return &arrayListIterator{elements: a.elements, current: 0}
}

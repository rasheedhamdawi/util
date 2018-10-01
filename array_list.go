package util

type arrayListIterator struct {
	current  int
	elements Elements
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
	Elements
}

// Size ...
func (a *ArrayList) Size() int {
	return len(a.Elements)
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
	if a.Elements == nil {
		a.Elements = make(map[int]Element)
	}

	// add a new element to the end of the list
	if a.Elements[a.Size()] = e; a.Elements[a.Size()-1] == e {
		return true
	}

	return false
}

// Contains ...
func (a *ArrayList) Contains(e Element) bool {

	for _, item := range a.Elements {
		if item == e {
			return true
		}
	}

	return false
}

// Get ...
func (a *ArrayList) Get(i int) Element {
	return a.Elements[i]
}

// IndexOf ...
func (a *ArrayList) IndexOf(e Element) int {

	for index, item := range a.Elements {
		if item == e {
			return index
		}
	}
	return -1
}

// Clear ...
func (a *ArrayList) Clear() {
	a.Elements = make(map[int]Element)
}

// Remove ...
func (a *ArrayList) Remove(e Element) error {
	delete(a.Elements, a.IndexOf(e))

	return nil
}

// Set ...
func (a *ArrayList) Set(index int, e Element) error {
	a.Elements[index] = e

	return nil
}

// AddAt ...
func (a *ArrayList) AddAt(index int, e Element) error {

	return nil
}

// ForEach ...
func (a *ArrayList) ForEach(cb func(int, Element)) {
	for i, v := range a.Elements {
		cb(i, v)
	}
}

// Iterator ...
func (a *ArrayList) Iterator() Iterator {
	return &arrayListIterator{elements: a.Elements, current: 0}
}

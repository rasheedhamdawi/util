package util

// ArrayList ...
type ArrayList struct {
	Elements map[int]Element
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

package util

import (
	"errors"
)

type hashSetIterator struct {
	items   []Element
	current int
}

func (s *hashSetIterator) Next() Element {
	defer func() {
		s.current++
	}()

	return s.items[s.current]
}

func (s *hashSetIterator) HasNext() bool {
	if s.current >= len(s.items) {
		return false
	}

	return true
}

type items map[Element]bool

// HashSet ...
type HashSet struct {
	hashSetIterator
	items
}

// Size ...
func (s *HashSet) Size() int {
	return len(s.items)
}

// IsEmpty ...
func (s *HashSet) IsEmpty() bool {
	if s.Size() > 0 {
		return false
	}

	return true
}

// Add ...
func (s *HashSet) Add(e Element) bool {
	// check if the list is empty make new one
	if s.items == nil {
		s.items = make(items)
	}

	if s.items[e] = true; s.items[e] {
		return true
	}

	return false
}

// Contains ...
func (s *HashSet) Contains(e Element) bool {
	for item := range s.items {
		if item == e {
			return true
		}
	}

	return false
}

// Clear ...
func (s *HashSet) Clear() {
	s.items = make(items)
}

// Remove ...
func (s *HashSet) Remove(e Element) error {
	delete(s.items, e)

	if s.items[e] {
		return errors.New("can't remove element")
	}

	return nil
}

// ForEach ...
func (s *HashSet) ForEach(cb func(int, Element)) {
	for i, v := range s.ToArray() {
		cb(i, v)
	}
}

// ToArray ...
func (s *HashSet) ToArray() []Element {
	var elements []Element

	for item := range s.items {
		elements = append(elements, item)
	}

	return elements
}

// Iterator ...
func (s *HashSet) Iterator() Iterator {
	return &hashSetIterator{items: s.ToArray(), current: 0}
}

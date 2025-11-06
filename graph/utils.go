package graph

// Source - https://stackoverflow.com/questions/62369223/the-equivalent-of-indexof
// Posted by thwd
// Retrieved 11/4/2025, License - CC-BY-SA 4.0

// IndexOf returns the first index of needle in haystack
// or -1 if needle is not in haystack.

// Set is a collection of unique elements
type Set struct {
	elements map[int]struct{}
}

// NewSet creates a new set
func NewSet() *Set {
	return &Set{
		elements: make(map[int]struct{}),
	}
}

// Add inserts an element into the set
func (s *Set) Add(value int) {
	s.elements[value] = struct{}{}
}

// Remove deletes an element from the set
func (s *Set) Remove(value int) {
	delete(s.elements, value)
}

// Contains checks if an element is in the set
func (s *Set) Contains(value int) bool {
	_, found := s.elements[value]
	return found
}

// Size returns the number of elements in the set
func (s *Set) Size() int {
	return len(s.elements)
}

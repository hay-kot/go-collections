// Package set implements a non-concurrent set.
package set

type Set[T comparable] struct {
	members map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{members: make(map[T]struct{})}
}

// Factory returns a new set containing the given values.
func Factory[T comparable](values ...T) *Set[T] {
	s := &Set[T]{
		members: make(map[T]struct{}, len(values)),
	}

	for _, v := range values {
		s.Add(v)
	}
	return s
}

// Add adds the given value to the set.
func (s *Set[T]) Add(v T) {
	s.members[v] = struct{}{}
}

// Remove removes the given value from the set.
func (s *Set[T]) Remove(v T) {
	delete(s.members, v)
}

// Len returns the number of values in the set.
func (s *Set[T]) Len() int {
	return len(s.members)
}

// IsEmpty returns true if the set is empty.
func (s *Set[T]) IsEmpty() bool {
	return len(s.members) == 0
}

// Clear removes all values from the set.
func (s *Set[T]) Clear() {
	s.members = make(map[T]struct{})
}

// Contains returns true if the set contains the given value.
func (s *Set[T]) Contains(v T) bool {
	_, ok := s.members[v]
	return ok
}

// ToSlice returns a slice of the set's values.
func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.members))
	for v := range s.members {
		slice = append(slice, v)
	}
	return slice
}

// Difference returns a new set containing the values that are in the
// receiver but not in the given set.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	diff := New[T]()
	for v := range s.members {
		if !other.Contains(v) {
			diff.Add(v)
		}
	}
	return diff
}

// Intersection returns a new set containing the values that are in both
// the receiver and the given set.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	intersection := New[T]()
	for v := range s.members {
		if other.Contains(v) {
			intersection.Add(v)
		}
	}
	return intersection
}

// Union returns a new set containing the values that are in either the
// receiver or the given set.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	union := New[T]()
	for v := range s.members {
		union.Add(v)
	}
	for v := range other.members {
		union.Add(v)
	}
	return union
}

// IsSubsetOf returns true if the receiver is a subset of the given set.
func (s *Set[T]) IsSubsetOf(other *Set[T]) bool {
	if s.Len() > other.Len() {
		return false
	}

	for v := range s.members {
		if !other.Contains(v) {
			return false
		}
	}

	return true
}

// Do calls the given function for each value in the set.
func (s *Set[T]) Do(f func(T)) {
	for v := range s.members {
		f(v)
	}
}

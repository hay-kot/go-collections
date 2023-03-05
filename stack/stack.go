// Package stack implements a non-concurrent and concurrent stack.
package stack

import "errors"

var (
	ErrEmptyStack = errors.New("stack is empty")
)

// Stack is a non-concurrent stack implementation.
type Stack[T any] struct {
	members []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) zero() T {
	var v T
	return v
}

func (s *Stack[T]) Push(v T) {
	s.members = append(s.members, v)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return s.zero(), ErrEmptyStack
	}

	v := s.members[len(s.members)-1]
	s.members = s.members[:len(s.members)-1]
	return v, nil
}

func (s *Stack[T]) Len() int {
	return len(s.members)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.members) == 0
}

func (s *Stack[T]) Clear() {
	s.members = nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return s.zero(), ErrEmptyStack
	}
	return s.members[len(s.members)-1], nil
}

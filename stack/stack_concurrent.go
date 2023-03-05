package stack

import "sync"

// ConcurrentStack is a ConcurrentStack implementation by wrapping the Stack within a mutex.
type ConcurrentStack[T any] struct {
	mu sync.Mutex
	Stack[T]
}

func NewConcurrentStack[T any]() *ConcurrentStack[T] {
	return &ConcurrentStack[T]{}
}

func (s *ConcurrentStack[T]) Push(v T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Stack.Push(v)
}

func (s *ConcurrentStack[T]) Pop() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.Stack.Pop()
}

func (cs *ConcurrentStack[T]) Clear() {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.Stack.Clear()
}

func (cs *ConcurrentStack[T]) Len() int {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	return cs.Stack.Len()
}

func (cs *ConcurrentStack[T]) IsEmpty() bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	return cs.Stack.IsEmpty()
}

func (cs *ConcurrentStack[T]) Peek() (T, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	return cs.Stack.Peek()
}

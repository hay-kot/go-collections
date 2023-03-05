package stack_test

import (
	"errors"
	"testing"

	"github.com/hay-kot/collections/stack"
)

// stacker is a stack interface for testing purposes.
type stacker[T any] interface {
	Push(v T)
	Pop() (T, error)
	Len() int
	IsEmpty() bool
	Clear()
	Peek() (T, error)
}

func TestStack_InterfaceImplementation(t *testing.T) {
	stacks := []stacker[string]{
		stack.NewStack[string](),
		stack.NewConcurrentStack[string](),
	}

	for i := range stacks {
		stk := stacks[i]

		vals := []string{
			"a",
			"b",
			"c",
		}

		for _, v := range vals {
			stk.Push(v)
		}

		if stk.Len() != len(vals) {
			t.Errorf("stack.Len() = %d, want %d", stk.Len(), len(vals))
		}

		for i := len(vals) - 1; i >= 0; i-- {
			v, err := stk.Pop()
			if err != nil {
				t.Errorf("stack.Pop() = %v, want %v", err, nil)
			}

			if v != vals[i] {
				t.Errorf("stack.Pop() = %v, want %v", v, vals[i])
			}
		}

		if !stk.IsEmpty() {
			t.Errorf("stack.IsEmpty() = %v, want %v", stk.IsEmpty(), true)
		}

		_, err := stk.Pop()
		if !errors.Is(err, stack.ErrEmptyStack) {
			t.Errorf("stack.Pop() = %v, want %v", err, stack.ErrEmptyStack)
		}
	}
}

func TestStack_Clear(t *testing.T) {
	stacks := []stacker[string]{
		stack.NewStack[string](),
		stack.NewConcurrentStack[string](),
	}

	for i := range stacks {
		stack := stacks[i]

		vals := []string{"a", "b", "c"}

		for _, v := range vals {
			stack.Push(v)
		}

		stack.Clear()

		if !stack.IsEmpty() {
			t.Errorf("stack.IsEmpty() = %v, want %v", stack.IsEmpty(), true)
		}

		if stack.Len() != 0 {
			t.Errorf("stack.Len() = %d, want %d", stack.Len(), 0)
		}
	}
}

func TestStack_Peak(t *testing.T) {
	stacks := []stacker[string]{
		stack.NewStack[string](),
		stack.NewConcurrentStack[string](),
	}

	for i := range stacks {
		stk := stacks[i]

		vals := []string{"a", "b", "c"}

		for _, v := range vals {
			stk.Push(v)
		}

		v, err := stk.Peek()
		if err != nil {
			t.Errorf("stack.Peek() = %v, want %v", err, nil)
		}

		if v != vals[len(vals)-1] {
			t.Errorf("stack.Peek() = %v, want %v", v, vals[len(vals)-1])
		}

		if stk.Len() != len(vals) {
			t.Errorf("stack.Len() = %d, want %d", stk.Len(), len(vals))
		}

		stk.Clear()

		_, err = stk.Peek()
		if !errors.Is(err, stack.ErrEmptyStack) {
			t.Errorf("stack.Peek() = %v, want %v", err, stack.ErrEmptyStack)
		}
	}
}

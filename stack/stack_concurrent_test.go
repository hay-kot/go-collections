package stack_test

import (
	"fmt"
	"sync"
	"testing"

	collections "github.com/hay-kot/collections/stack"
)

const ROUTINES = 5

func TestConcurrentStack_Push(t *testing.T) {
	stack := collections.NewConcurrentStack[string]()

	wg := sync.WaitGroup{}
	wg.Add(ROUTINES)
	for i := 0; i < ROUTINES; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				stack.Push(fmt.Sprintf("routine %d, iteration %d", i, j))
			}
		}(i)
	}

	wg.Wait()

	if stack.Len() != ROUTINES*100 {
		t.Errorf("stack.Len() = %d, want %d", stack.Len(), ROUTINES*100)
	}
}

func TestConcurrentStack_Pop(t *testing.T) {
	stack := collections.NewConcurrentStack[string]()

	wg := sync.WaitGroup{}
	wg.Add(ROUTINES)

	for i := 0; i < ROUTINES; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				stack.Push(fmt.Sprintf("routine %d, iteration %d", i, j))
				_, err := stack.Pop()
				if err != nil {
					t.Errorf("stack.Pop() = %v, want %v", err, nil)
				}
			}
		}(i)
	}

	wg.Wait()
	if stack.Len() != 0 {
		t.Errorf("stack.Len() = %d, want %d", stack.Len(), ROUTINES*100)
	}
}

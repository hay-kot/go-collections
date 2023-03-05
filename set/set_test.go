package set_test

import (
	"testing"

	"github.com/hay-kot/collections/set"
)

func Test_ToSlice(t *testing.T) {
	s := set.New[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	slice := s.ToSlice()
	if len(slice) != 3 {
		t.Errorf("Expected slice to have length 3, got %d", len(slice))
	}

	for _, v := range slice {
		if !s.Contains(v) {
			t.Errorf("Expected slice to contain %d", v)
		}
	}
}

func TestSet_Difference(t *testing.T) {
	s1 := set.Factory(1, 2, 3, 4, 5)
	s2 := set.Factory(1, 2, 3)

	s3 := s1.Difference(s2)

	if s3.Len() != 2 {
		t.Errorf("Expected difference to have length 2, got %d", s3.Len())
	}

	for _, v := range s3.ToSlice() {
		if s2.Contains(v) {
			t.Errorf("Expected difference to not contain %d", v)
		}
	}
}

func TestSet_Intersection(t *testing.T) {
	s1 := set.Factory(1, 2, 3, 4, 5)
	s2 := set.Factory(1, 2, 3)

	s3 := s1.Intersection(s2)

	if s3.Len() != 3 {
		t.Errorf("Expected intersection to have length 3, got %d", s3.Len())
	}

	for _, v := range s3.ToSlice() {
		if !s2.Contains(v) {
			t.Errorf("Expected intersection to contain %d", v)
		}
	}
}

func TestSet_Union(t *testing.T) {
	s1 := set.Factory(1, 2, 3, 4, 5)
	s2 := set.Factory(1, 2, 3)

	s3 := s1.Union(s2)

	if s3.Len() != 5 {
		t.Errorf("Expected union to have length 5, got %d", s3.Len())
	}

	for _, v := range s3.ToSlice() {
		if !s1.Contains(v) {
			t.Errorf("Expected union to contain %d", v)
		}
	}
}

func TestSet_IsSubset(t *testing.T) {
	s1 := set.Factory(1, 2, 3, 4, 5)
	s2 := set.Factory(1, 2, 3)
	s3 := set.Factory(1, 2, 3, 4, 5, 6)

	if !s2.IsSubsetOf(s1) {
		t.Errorf("Expected s2 to be a subset of s1")
	}

	if s1.IsSubsetOf(s2) {
		t.Errorf("Expected s1 not to be a subset of s2")
	}

	if !s1.IsSubsetOf(s3) {
		t.Errorf("Expected s1 to be a subset of s3")
	}
}

func TestSet_Do(t *testing.T) {
	count := 0

	s := set.Factory(1, 2, 3, 4, 5)
	s.Do(func(v int) {
		if !s.Contains(v) {
			t.Errorf("Expected set to contain %d", v)
		}

		count++
	})

	if count != s.Len() {
		t.Error("Expected count to equal set length")
	}
}

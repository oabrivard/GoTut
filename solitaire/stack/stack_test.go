/*
Copyright 2022 Olivier Abrivard

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.
*/

package stack

import "testing"

func TestNegativeSize(t *testing.T) {
	defer func() { _ = recover() }() // turn off the panic

	_ = New[int](-1)

	// Never reaches here if New() panics.
	t.Errorf("New() with negative size should panic")
}

func TestEmpty(t *testing.T) {
	defer func() { _ = recover() }() // turn off the panic

	s := New[int](0)
	s.Pop()

	// Never reaches here if Pop() panics.
	t.Errorf("Pop() on empty stack should panic")
}

func TestFull(t *testing.T) {
	defer func() { _ = recover() }() // turn off the panic

	s := New[int](1)
	s.Push(1)
	s.Push(2)

	// Never reaches here if Push() panics.
	t.Errorf("Push() on full stack should panic")
}

func TestHasElements(t *testing.T) {
	s := New[int](1)

	if s.HasElement() {
		t.Errorf("HasElement() should return false for empty stack")
	}

	s.Push(1)

	if !s.HasElement() {
		t.Errorf("HasElement() should return true for non empty stack")
	}
}

func TestCount(t *testing.T) {
	s := New[int](1)

	if s.Count() != 0 {
		t.Errorf("Count() should be 0 for empty stack")
	}

	s.Push(1)

	if s.Count() != 1 {
		t.Errorf("Count() should be 1 for stack with as single element")
	}
}

func TestPushPop(t *testing.T) {
	s1 := New[int](3)
	s1.Push(1)
	s1.Push(2)
	s1.Push(3)

	var expectedValues = []int{3, 2, 1}

	for _, expected := range expectedValues {
		got := s1.Pop()

		if expected != got {
			t.Errorf("got %d, want %d", got, expected)
		}
	}
}

func TestReversed(t *testing.T) {
	s1 := New[int](3)
	s1.Push(1)
	s1.Push(2)
	s1.Push(3)

	s2 := s1.Reversed()

	if s1.Count() != 0 {
		t.Errorf("Reversed() should emty the stack it is called upon")
	}

	var expectedValues = []int{1, 2, 3}

	for _, expected := range expectedValues {
		got := s2.Pop()

		if expected != got {
			t.Errorf("got %d, want %d", got, expected)
		}
	}
}

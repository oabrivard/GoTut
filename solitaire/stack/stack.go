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

// stack head == 0 if stack is empty.
// stack head == size if stack is full.
type Stack[T any] struct {
	size int
	head int
	data []T
}

// New creates an empty stack that can contain a maximum of size elements.
func New[T any](size int) Stack[T] {
	if size < 0 {
		panic("stack size must be >= 0")
	}

	result := Stack[T]{
		size,
		0,
		nil,
	}

	if size > 0 {
		result.data = make([]T, size)
	}

	return result
}

// Push pushes the value v on top of stack s.
func (s *Stack[T]) Push(v T) {
	if s.head == s.size {
		panic("called Push() on a full stack")
	}

	s.data[s.head] = v
	s.head++
}

// Pop removes the value v from the top of stack s and returns it.
func (s *Stack[T]) Pop() T {
	if !s.HasElement() {
		panic("called Pop() on an empty stack")
	}

	s.head--
	return s.data[s.head]
}

// HasElement returns true if the stack s contains at least one element.
func (s *Stack[T]) HasElement() bool {
	return s.head > 0
}

// Count returns the number of elements contained in s.
func (s *Stack[T]) Count() int {
	return s.head
}

// Reversed returns a new stack built from stack s, with its elements in reverse order.
// s will be emptied by this operation.
func (s *Stack[T]) Reversed() Stack[T] {
	result := Stack[T]{
		s.size,
		0,
		make([]T, s.size),
	}

	for s.HasElement() {
		result.Push(s.Pop())
	}

	return result
}

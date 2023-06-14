package main

import "fmt"

type Stack[T any] struct {
	top    *node[T]
	length int
}
type node[T any] struct {
	value T
	prev  *node[T]
}

// Create a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil, 0}
}

// Return the number of items in the stack
func (this *Stack[T]) Len() int {
	return this.length
}

// View the top item on the stack
func (this *Stack[T]) Peek() T {
	if this.length == 0 {
		return *new(T)
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack[T]) Pop() T {
	if this.length == 0 {
		return *new(T)
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack[T]) Push(value T) {
	n := &node[T]{value, this.top}
	this.top = n
	this.length++
}

type Thing struct {
	a string
	b int
}

func main() {
	s1 := NewStack[string]()
	s1.Push("one")
	s1.Push("two")
	fmt.Println(s1.Pop())
	fmt.Println(s1.Pop())
	fmt.Println(s1.Pop())

	s2 := NewStack[*Thing]()
	s2.Push(&Thing{a: "hey", b: 1})
	s2.Push(&Thing{a: "ho", b: 2})
	fmt.Println(s2.Pop())
	fmt.Println(s2.Pop())
	fmt.Println(s2.Pop())
}

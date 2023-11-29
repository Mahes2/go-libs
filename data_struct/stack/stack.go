package stack

type Stack[T comparable] struct {
	items []T
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

func (stack *Stack[T]) Push(element T) {
	stack.items = append(stack.items, element)
}

func (stack *Stack[T]) Peek() T {
	var lastElement T

	if len(stack.items) > 0 {
		lastElement = stack.items[len(stack.items)-1]
	}

	return lastElement
}

func (stack *Stack[T]) Pop() T {
	lastElement := stack.Peek()

	stack.items = stack.items[:len(stack.items)-1]

	return lastElement
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.items) == 0
}

func (stack *Stack[T]) Size() int {
	return len(stack.items)
}

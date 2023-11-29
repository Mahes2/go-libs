package stack

import "testing"

func TestNewStack(t *testing.T) {
	stack := NewStack[int]()
	if stack == nil {
		t.Error("stack is nil")
	}
}

func TestPush(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Size() != 3 {
		t.Error("stack size is not 3")
	}
}

func TestPop(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)

	if stack.Pop() != 1 {
		t.Error("stack top is not 1")
	}

	if !stack.IsEmpty() {
		t.Error("stack is not empty")
	}
}

func TestPeek(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Peek() != 3 {
		t.Error("stack top is not 3")
	}

	if stack.Size() != 3 {
		t.Error("stack size is not 3")
	}
}

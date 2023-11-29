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

func TestClear(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Clear()

	if !stack.IsEmpty() {
		t.Error("stack is not empty")
	}
}

func BenchmarkStackPush100(b *testing.B) {
	stack := NewStack[int]()
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			stack.Push(j)
		}
	}
}

func BenchmarkStackPush10000(b *testing.B) {
	stack := NewStack[int]()
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			stack.Push(j)
		}
	}
}

func BenchmarkStackPush1000000(b *testing.B) {
	stack := NewStack[int]()
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000000; j++ {
			stack.Push(j)
		}
	}
}

func BenchmarkStackPop100(b *testing.B) {
	stack := NewStack[int]()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for j := 0; j < 100; j++ {
			stack.Push(j)
		}
		b.StartTimer()
		for j := 0; j < 100; j++ {
			stack.Pop()
		}
	}
}

func BenchmarkStackPop10000(b *testing.B) {
	stack := NewStack[int]()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for j := 0; j < 10000; j++ {
			stack.Push(j)
		}
		b.StartTimer()
		for j := 0; j < 10000; j++ {
			stack.Pop()
		}
	}
}

func BenchmarkStackPop1000000(b *testing.B) {
	stack := NewStack[int]()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for j := 0; j < 1000000; j++ {
			stack.Push(j)
		}
		b.StartTimer()
		for j := 0; j < 1000000; j++ {
			stack.Pop()
		}
	}
}

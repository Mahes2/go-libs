package stack

import (
	"fmt"
	"go-libs/data_struct/stack"
)

func main() {
	stack := stack.NewStack[int]()
	if stack.IsEmpty() {
		fmt.Println("stack is empty")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("stack size is: ", stack.Size())

	fmt.Println("stack last element is: ", stack.Peek())

	_ = stack.Pop()
	fmt.Println("stack last element is: ", stack.Peek())

	stack.Clear()
	if stack.IsEmpty() {
		fmt.Println("stack is empty")
	}
}

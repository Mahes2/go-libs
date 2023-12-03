package errortracer

import "testing"

func TestInitNewStackSize(t *testing.T) {
	newStackSize := 30

	Init(Option{
		StackSize: newStackSize,
	})

	if stackSize != newStackSize {
		t.Errorf("Expected stack size is %d but got %d", newStackSize, stackSize)
	}
}

func TestInitInvalidStackSize(t *testing.T) {
	newStackSize := -3

	Init(Option{
		StackSize: newStackSize,
	})

	if stackSize != DEFAULT_STACK_SIZE {
		t.Errorf("Expected default stack size (%d) but got %d", DEFAULT_STACK_SIZE, stackSize)
	}
}

func TestSelfInit(t *testing.T) {
	selfInit()

	if stackSize != DEFAULT_STACK_SIZE {
		t.Errorf("Expected default stack size (%d) but got %d", DEFAULT_STACK_SIZE, stackSize)
	}
}

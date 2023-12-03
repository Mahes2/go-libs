package errortracer

const (
	DEFAULT_STACK_SIZE = 10
)

var (
	stackSize int
)

type Option struct {
	StackSize int
}

func Init(option Option) {
	if option.StackSize < DEFAULT_STACK_SIZE {
		option.StackSize = DEFAULT_STACK_SIZE
	}

	stackSize = option.StackSize
}

func selfInit() {
	stackSize = DEFAULT_STACK_SIZE
}

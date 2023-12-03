package main

import (
	"fmt"

	errortracer "github.com/Mahes2/go-libs/tracer/grpc_error_tracer"
	"google.golang.org/grpc/codes"
)

func main() {
	errortracer.Init(errortracer.Option{
		StackSize: 0,
	})

	err := testFunction1()
	newError := errortracer.WrapWithDataAndLog(err, map[string]interface{}{
		"request": "test",
	})

	_ = newError.Error()
}

func testFunction1() error {
	err := testFunction2()
	return errortracer.Wrap(err)
}

func testFunction2() error {
	err := testFunction3()
	return errortracer.WrapWithData(err, map[string]interface{}{
		"response": "test2",
	})
}

func testFunction3() error {
	err := fmt.Errorf("this a sample exception")
	return errortracer.NewError(codes.InvalidArgument, err.Error(), "internal server error")
}

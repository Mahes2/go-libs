package main

import (
	"fmt"

	errortracer "github.com/Mahes2/go-libs/tracer/error_tracer"
)

func main() {
	err := testFunction1()
	err = errortracer.AddData(err, map[string]interface{}{
		"request": "test",
	})

	fmt.Println(errortracer.Print(err))
}

func testFunction1() error {
	err := testFunction2()
	return err
}

func testFunction2() error {
	err := testFunction3()
	return errortracer.AddData(err, map[string]interface{}{
		"response": "test2",
	})
}

func testFunction3() error {
	err := errortracer.NewError("this a sample exception", "internal server error")
	return err
}

package errortracer

import (
	"fmt"
	"testing"
)

func TestNewError(t *testing.T) {
	originalErr := "this is a sample exception"
	customErr := "internal server error"

	newError := NewError(originalErr, customErr)

	_, ok := newError.(*errorTracer)
	if !ok {
		t.Errorf("new error is not errorTracer")
	}

	if newError.Error() != customErr {
		t.Errorf("got error: %s, expected: %s", newError.Error(), customErr)
	}
}

func TestNewErrorWithData(t *testing.T) {
	originalErr := "this is a sample exception"
	additionalData := map[string]interface{}{
		"name": "go-libs",
	}

	newError := NewErrorWithData(originalErr, "", additionalData)

	errTracer, ok := newError.(*errorTracer)
	if !ok {
		t.Errorf("new error is not errorTracer")
	}

	if newError.Error() != originalErr {
		t.Errorf("got error: %s, expected: %s", newError.Error(), originalErr)
	}

	if len(errTracer.additionalData) != len(additionalData) {
		t.Errorf("got data length is: %d, want: %d", len(errTracer.additionalData), len(additionalData))
	}
}

func TestWrapError(t *testing.T) {
	err := fmt.Errorf("this is a sample exception")

	newError := Wrap(err, "")

	_, ok := newError.(*errorTracer)
	if !ok {
		t.Errorf("new error is not errorTracer")
	}

	if newError.Error() != err.Error() {
		t.Errorf("got error: %s, expected: %s", newError.Error(), err.Error())
	}
}

func TestWrapErrorWithData(t *testing.T) {
	err := fmt.Errorf("this is a sample exception")
	additionalData := map[string]interface{}{
		"name": "go-libs",
	}

	newError := WrapWithData(err, "", additionalData)

	errTracer, ok := newError.(*errorTracer)
	if !ok {
		t.Errorf("new error is not errorTracer")
	}

	if newError.Error() != err.Error() {
		t.Errorf("got error: %s, expected: %s", newError.Error(), err.Error())
	}

	if len(errTracer.additionalData) != len(additionalData) {
		t.Errorf("got data length is: %d, want: %d", len(errTracer.additionalData), len(additionalData))
	}
}

func TestAddData(t *testing.T) {
	err := &errorTracer{
		originalMessage: "this is a sample exception",
	}
	additionalData := map[string]interface{}{
		"name": "go-libs",
	}

	newError := AddData(err, additionalData)

	errTracer, ok := newError.(*errorTracer)
	if !ok {
		t.Errorf("new error is not errorTracer")
	}

	if len(errTracer.additionalData) != 1 {
		t.Errorf("got data length is: %d, want: %d", len(errTracer.additionalData), len(additionalData))
	}
}

func TestPrint(t *testing.T) {
	originalErr := "this is a sample exception"
	customErr := "internal server error"
	additionalData := map[string]interface{}{
		"name": "go-libs",
	}

	newError := NewErrorWithData(originalErr, customErr, additionalData)
	output := Print(newError)

	if output == "" {
		t.Errorf("got empty output")
	}
}

func BenchmarkWrapError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := fmt.Errorf("this is a sample exception")
		_ = Wrap(err, "")
	}
}

func BenchmarkWrapAndPrintError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := fmt.Errorf("this is a sample exception")
		err = Wrap(err, "")
		_ = Print(err)
	}
}

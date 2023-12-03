package errortracer

import (
	"fmt"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNewError(t *testing.T) {
	originalErr := status.Error(codes.InvalidArgument, "this is a sample exception")
	customErr := "internal server error"
	newError := NewError(status.Code(originalErr), originalErr.Error(), customErr)
	if newError.Error() != customErr {
		t.Errorf("new error is not as expected: %s", customErr)
	}
}

func TestNewErrorWithData(t *testing.T) {
	originalErr := fmt.Errorf("this is a sample exception")
	additionalData := map[string]interface{}{
		"name": "go-libs",
	}
	newError := NewErrorWithData(status.Code(originalErr), originalErr.Error(), "", additionalData)
	if newError.Error() != originalErr.Error() {
		t.Errorf("new error is not as expected: %s", originalErr.Error())
	}
}

func TestWrapError(t *testing.T) {
	err := fmt.Errorf("this is a sample exception")
	newError := Wrap(err)
	if newError.Error() != err.Error() {
		t.Errorf("new error is not as expected: %s", err.Error())
	}
}

func TestWrapErrorWithAdditionalData(t *testing.T) {
	err := fmt.Errorf("this is a sample exception")
	newError := WrapWithData(err, map[string]interface{}{
		"name": "go-libs",
	})
	if newError.Error() != err.Error() {
		t.Errorf("new error is not as expected: %s", err.Error())
	}
}

func TestGrowStackTraceSize(t *testing.T) {
	err := fmt.Errorf("this is a sample exception")
	newError := Wrap(err)
	newError = WrapWithData(newError, map[string]interface{}{
		"name": "go-libs",
	})
	newError = Wrap(newError)
	newError = WrapWithData(newError, map[string]interface{}{
		"name": "go-libs",
	})
	newError = Wrap(newError)
	newError = WrapWithData(newError, map[string]interface{}{
		"name": "go-libs",
	})
	newError = Wrap(newError)
	newError = WrapWithData(newError, map[string]interface{}{
		"name": "go-libs",
	})
	newError = Wrap(newError)
	newError = WrapWithData(newError, map[string]interface{}{
		"name": "go-libs",
	})
	newError = Wrap(newError)
	newError = WrapWithData(newError, map[string]interface{}{
		"name": "go-libs",
	})
	newError = Wrap(newError)
	newError = WrapWithData(newError, map[string]interface{}{
		"name": "go-libs",
	})
	if newError.Error() != err.Error() {
		t.Errorf("new error is not as expected: %s", err.Error())
	}
}

func BenchmarkWrapError100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := fmt.Errorf("this is a sample exception")
		for j := 0; j < 100; j++ {
			err = Wrap(err)
		}
	}
}

func BenchmarkWrapError10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := fmt.Errorf("this is a sample exception")
		for j := 0; j < 10000; j++ {
			err = Wrap(err)
		}
	}
}

func BenchmarkWrapError1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := fmt.Errorf("this is a sample exception")
		for j := 0; j < 1000000; j++ {
			err = Wrap(err)
		}
	}
}
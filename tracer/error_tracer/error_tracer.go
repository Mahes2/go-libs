package errortracer

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

type errorTracer struct {
	originalError  error
	customError    error
	additionalData map[string]interface{}
	stackTrace     []string
	traceCount     int
}

func (errTracer *errorTracer) Error() string {
	errTracer.log()

	if errTracer.customError != nil {
		return errTracer.customError.Error()
	}

	return errTracer.originalError.Error()
}

func NewError(originalError, customError error) error {
	return newErrorTracer(originalError, customError, nil)
}

func NewErrorWithData(originalError, customError error, additionalData map[string]interface{}) error {
	return newErrorTracer(originalError, customError, additionalData)
}

func Wrap(err error) error {
	errTracer, ok := err.(*errorTracer)
	if !ok {
		return newErrorTracer(err, nil, nil)
	}

	errTracer.addTrace(4)

	return errTracer
}

func WrapWithData(err error, additionalData map[string]interface{}) error {
	errTracer, ok := err.(*errorTracer)
	if !ok {
		return newErrorTracer(err, nil, additionalData)
	}

	errTracer.addMoreData(additionalData)
	errTracer.addTrace(4)

	return errTracer
}

func newErrorTracer(originalError, customError error, additionalData map[string]interface{}) *errorTracer {
	errTracer := &errorTracer{
		originalError: originalError,
		customError:   customError,
	}

	errTracer.addMoreData(additionalData)
	errTracer.addTrace(5)

	return errTracer
}

func (errTracer *errorTracer) addMoreData(additionalData map[string]interface{}) {
	if additionalData == nil {
		return
	}

	if errTracer.additionalData == nil {
		errTracer.additionalData = make(map[string]interface{})
	}

	for k, v := range additionalData {
		errTracer.additionalData[k] = v
	}
}

func (errTracer *errorTracer) addTrace(skip int) {
	if errTracer.stackTrace == nil {
		errTracer.createStackTrace()
	}

	if errTracer.traceCount == len(errTracer.stackTrace) {
		errTracer.growStackTrace()
	}

	errTracer.stackTrace[errTracer.traceCount] = getCallerDetail(skip)
	errTracer.traceCount++
}

func (errTracer *errorTracer) createStackTrace() {
	if stackSize < DEFAULT_STACK_SIZE {
		selfInit()
	}

	errTracer.stackTrace = make([]string, stackSize)
}

func (errTracer *errorTracer) growStackTrace() {
	newTraces := make([]string, 2*len(errTracer.stackTrace))
	copy(newTraces, errTracer.stackTrace)
	errTracer.stackTrace = newTraces
}

func getCallerDetail(skip int) string {
	pc := make([]uintptr, 10)
	runtime.Callers(skip, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	caller := fmt.Sprintf("%s:%d %s", file, line, f.Name())

	return caller
}

func (errTracer *errorTracer) log() {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Original Error: %s\n", errTracer.getOriginalError().Error()))
	sb.WriteString(fmt.Sprintf("Custom Error: %s\n", errTracer.getCustomError().Error()))

	sb.WriteString(strings.Join(errTracer.stackTrace, "\n"))

	if errTracer.additionalData != nil {
		sb.WriteString("\nAdditional Data:")
		for key, value := range errTracer.additionalData {
			jsonStr, _ := json.Marshal(value)
			sb.WriteString(fmt.Sprintf("\n%s: %v", key, string(jsonStr)))
		}
	}

	fmt.Println(sb.String())
}

func (errTracer *errorTracer) getOriginalError() error {
	return errTracer.originalError
}

func (errTracer *errorTracer) getCustomError() error {
	if errTracer.customError == nil {
		return fmt.Errorf("")
	}

	return errTracer.customError
}

package errortracer

import (
	"encoding/json"
	"runtime"
	"strconv"
	"strings"
)

type errorTracer struct {
	originalError  string
	userMessage    string
	additionalData map[string]any
	stackTrace     []uintptr
}

func (errTracer *errorTracer) Error() string {
	if errTracer.userMessage != "" {
		return errTracer.userMessage
	}

	return errTracer.originalError
}

func NewError(originalError, userMessage string) error {
	return newErrorTracer(originalError, userMessage, nil)
}

func NewErrorWithData(originalError, userMessage string, additionalData map[string]any) error {
	return newErrorTracer(originalError, userMessage, additionalData)
}

func Wrap(err error, userMessage string) error {
	if err == nil {
		return err
	}

	errTracer, ok := err.(*errorTracer)
	if !ok {
		return newErrorTracer(err.Error(), userMessage, nil)
	}

	errTracer.userMessage = userMessage
	return errTracer
}

func WrapWithData(err error, userMessage string, additionalData map[string]any) error {
	if err == nil {
		return err
	}

	errTracer, ok := err.(*errorTracer)
	if !ok {
		return newErrorTracer(err.Error(), userMessage, additionalData)
	}

	errTracer.userMessage = userMessage
	errTracer.addData(additionalData)
	return errTracer

}

func AddData(err error, additionalData map[string]any) error {
	if err == nil {
		return err
	}

	errTracer, ok := err.(*errorTracer)
	if !ok {
		return err
	}

	errTracer.addData(additionalData)
	return errTracer
}

func Print(err error) string {
	if err == nil {
		return ""
	}

	errTracer, ok := err.(*errorTracer)
	if !ok {
		return err.Error()
	}

	return errTracer.print()
}

func newErrorTracer(originalError, userMessage string, additionalData map[string]any) *errorTracer {
	errTracer := &errorTracer{
		originalError: originalError,
		userMessage:   userMessage,
	}

	errTracer.addData(additionalData)
	errTracer.addTrace()

	return errTracer
}

func (errTracer *errorTracer) addData(additionalData map[string]any) {
	if additionalData == nil {
		return
	}

	if errTracer.additionalData == nil {
		errTracer.additionalData = make(map[string]any)
	}

	for k, v := range additionalData {
		errTracer.additionalData[k] = v
	}
}

func (errTracer *errorTracer) addTrace() {
	errTracer.stackTrace = getCallerDetail()
}

func getCallerDetail() []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(5, pcs[:])
	return pcs[0:n]
}

func (errTracer *errorTracer) print() string {
	var sb strings.Builder

	sb.WriteString("Original Error: ")
	sb.WriteString(errTracer.originalError)
	sb.WriteString("\nUser Message: ")
	sb.WriteString(errTracer.userMessage)

	sb.WriteString("\n\nTraces: \n")
	for k := range errTracer.stackTrace {
		v := errTracer.stackTrace[k] - 1
		f := runtime.FuncForPC(v)
		file, line := f.FileLine(v)

		sb.WriteString(f.Name())
		sb.WriteString("\n\t")
		sb.WriteString(file)
		sb.WriteString(":")
		sb.WriteString(strconv.Itoa(line))
		sb.WriteString("\n")
	}

	sb.WriteString("\nAdditional Data: ")
	for key, value := range errTracer.additionalData {
		jsonStr, _ := json.Marshal(value)
		sb.WriteString("\n")
		sb.WriteString(key)
		sb.WriteString(": ")
		sb.WriteString(string(jsonStr))
	}

	return sb.String()
}

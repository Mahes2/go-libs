package errortracer

import (
	"encoding/json"
	"runtime"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type errorTracer struct {
	code            codes.Code
	originalMessage string
	userMessage     string
	additionalData  map[string]interface{}
	stackTrace      []uintptr
}

func (errTracer *errorTracer) Error() string {
	if errTracer.userMessage != "" {
		return errTracer.userMessage
	}

	return errTracer.originalMessage
}

func (errTracer *errorTracer) GRPCStatus() *status.Status {
	return status.New(errTracer.code, errTracer.Error())
}

func NewError(code codes.Code, originalMessage, userMessage string) error {
	return newErrorTracer(code, originalMessage, userMessage, nil)
}

func NewErrorWithData(code codes.Code, originalMessage, userMessage string, additionalData map[string]interface{}) error {
	return newErrorTracer(code, originalMessage, userMessage, additionalData)
}

func Wrap(err error, userMessage string) error {
	if err == nil {
		return err
	}

	errTracer, ok := err.(*errorTracer)
	if !ok {
		return newErrorTracer(status.Code(err), err.Error(), userMessage, nil)
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
		return newErrorTracer(status.Code(err), err.Error(), userMessage, additionalData)
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

func newErrorTracer(code codes.Code, originalMessage, userMessage string, additionalData map[string]interface{}) *errorTracer {
	errTracer := &errorTracer{
		code:            code,
		originalMessage: originalMessage,
		userMessage:     userMessage,
		stackTrace:      getCallerDetail(),
	}

	errTracer.addData(additionalData)

	return errTracer
}

func (errTracer *errorTracer) addData(additionalData map[string]interface{}) {
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

func getCallerDetail() []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(5, pcs[:])
	return pcs[0:n]
}

func (errTracer *errorTracer) print() string {
	var sb strings.Builder

	sb.WriteString("Status Code: ")
	sb.WriteString(errTracer.code.String())

	sb.WriteString("\nOriginal Error: ")
	sb.WriteString(errTracer.originalMessage)
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

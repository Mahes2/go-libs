package errortracer

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type errorTracer struct {
	code           codes.Code
	originalError  string
	customError    string
	additionalData map[string]interface{}
	stackTrace     []string
	traceCount     int
}

func (errTracer *errorTracer) Error() string {
	errTracer.log()

	if errTracer.customError != "" {
		return errTracer.customError
	}

	return errTracer.originalError
}

func (errTracer *errorTracer) GRPCStatus() *status.Status {
	return status.New(errTracer.code, errTracer.Error())
}

func NewError(code codes.Code, originalError, customError string) error {
	return newErrorTracer(code, originalError, customError, nil)
}

func NewErrorWithData(code codes.Code, originalError, customError string, additionalData map[string]interface{}) error {
	return newErrorTracer(code, originalError, customError, additionalData)
}

func Wrap(err error) error {
	errTracer, ok := err.(*errorTracer)
	if !ok {
		return newErrorTracer(status.Code(err), err.Error(), "", nil)
	}

	errTracer.addTrace(4)

	return errTracer
}

func WrapWithData(err error, additionalData map[string]interface{}) error {
	errTracer, ok := err.(*errorTracer)
	if !ok {
		return newErrorTracer(status.Code(err), err.Error(), "", additionalData)
	}

	errTracer.addMoreData(additionalData)
	errTracer.addTrace(4)

	return errTracer
}

func newErrorTracer(code codes.Code, originalError, customError string, additionalData map[string]interface{}) *errorTracer {
	errTracer := &errorTracer{
		code:          code,
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

	sb.WriteString(fmt.Sprintf("Status Code: %s\n", errTracer.getStatusCode()))

	sb.WriteString(fmt.Sprintf("Original Error: %s\n", errTracer.getOriginalError()))
	sb.WriteString(fmt.Sprintf("Custom Error: %s\n\n", errTracer.getCustomError()))

	sb.WriteString(fmt.Sprintf("Traces: \n%s\n", errTracer.getStackTraceList()))

	sb.WriteString(fmt.Sprintf("Additional Data: %s", errTracer.getAdditionalDataList()))

	fmt.Println(sb.String())
}

func (errTracer *errorTracer) getStatusCode() string {
	return errTracer.code.String()
}

func (errTracer *errorTracer) getOriginalError() string {
	return errTracer.originalError
}

func (errTracer *errorTracer) getCustomError() string {
	return errTracer.customError
}

func (errTracer *errorTracer) getStackTraceList() string {
	var str string
	for _, v := range errTracer.stackTrace {
		if v == "" {
			break
		}
		str += v + "\n"
	}

	return str
}

func (errTracer *errorTracer) getAdditionalDataList() string {
	var str string
	for key, value := range errTracer.additionalData {
		jsonStr, _ := json.Marshal(value)
		str += fmt.Sprintf("\n%s: %v", key, string(jsonStr))
	}

	return str
}

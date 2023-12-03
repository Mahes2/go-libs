package errortracer

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

type errorTracer struct {
	originalError  string
	customError    string
	additionalData map[string]interface{}
	stackTrace     []string
	traceCount     int
}

func (errTracer *errorTracer) Error() string {
	if errTracer.customError != "" {
		return errTracer.customError
	}

	return errTracer.originalError
}

func NewError(originalError, customError string) error {
	return newErrorTracer(originalError, customError, nil)
}

func NewErrorWithData(originalError, customError string, additionalData map[string]interface{}) error {
	return newErrorTracer(originalError, customError, additionalData)
}

func Wrap(err error) error {
	return wrap(err, 4)
}

func WrapAndLog(err error) error {
	errTracer := wrap(err, 5)
	errTracer.log()

	return errTracer
}

func WrapWithData(err error, additionalData map[string]interface{}) error {
	errTracer := wrap(err, 5)
	errTracer.addMoreData(additionalData)

	return errTracer
}

func WrapWithDataAndLog(err error, additionalData map[string]interface{}) error {
	errTracer := wrap(err, 5)
	errTracer.addMoreData(additionalData)
	errTracer.log()

	return errTracer
}

func newErrorTracer(originalError, customError string, additionalData map[string]interface{}) *errorTracer {
	errTracer := &errorTracer{
		originalError: originalError,
		customError:   customError,
	}

	errTracer.addMoreData(additionalData)
	errTracer.addTrace(5)

	return errTracer
}

func wrap(err error, skip int) *errorTracer {
	errTracer, ok := err.(*errorTracer)
	if !ok {
		return newErrorTracer(err.Error(), "", nil)
	}

	errTracer.addTrace(skip)

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

	sb.WriteString(fmt.Sprintf("Original Error: %s\n", errTracer.getOriginalError()))
	sb.WriteString(fmt.Sprintf("Custom Error: %s\n\n", errTracer.getCustomError()))

	sb.WriteString(fmt.Sprintf("Traces: \n%s\n", errTracer.getStackTraceList()))

	sb.WriteString(fmt.Sprintf("Additional Data: %s", errTracer.getAdditionalDataList()))

	fmt.Println(sb.String())
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

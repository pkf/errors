package errors

import (
	"runtime"
	"fmt"
)

type PreDefineCode int

func (e PreDefineCode) Int() int {
	return int(e)
}
func (e PreDefineCode) New(s interface{}) *Error {
	pcs := make([]uintptr, 32)
	count := runtime.Callers(2, pcs)
	e1 := &Error{
		Info:    fmt.Sprint(s),
		Code:    int(e),
		stackPC: pcs[:count],
	}
	switch raw := s.(type) {
	case error:
		e1.rawErr = raw
	default:
		e1.rawErr = e1
	}
	return e1
}
func (e PreDefineCode) Errorf(format string, a ...interface{}) *Error {
	pcs := make([]uintptr, 32)
	count := runtime.Callers(2, pcs)
	e1 := &Error{
		Info:    fmt.Sprintf(format, a...),
		Code:    int(e),
		stackPC: pcs[:count],
	}
	e1.rawErr = e1
	return e1
}

func (e PreDefineCode) Errorln(a ...interface{}) *Error {
	pcs := make([]uintptr, 32)
	count := runtime.Callers(2, pcs)

	e1 := &Error{
		Info:    fmt.Sprintln(a...),
		Code:    int(e),
		stackPC: pcs[:count],
	}
	e1.rawErr = e1
	return e1

}

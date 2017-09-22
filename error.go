package errors

import (
	"runtime"
	"fmt"
	"strings"
	"sync/atomic"
)

var debug int32= 0

func EnableDebug(){
	atomic.StoreInt32(&debug,1)
}
func DisableDebug(){
	atomic.StoreInt32(&debug,0)
}

type Error struct {
	Info string
	rawErr error     // 保存原始错误信息
	stackPC []uintptr // 保存函数调用栈指针
}
//func (e *Error)Error()string{
//
//	return fmt.Sprintf("%s:%d:%s",e.File,e.Line,e.Info)
//}

// CallStack get function call stack
func (e *Error) Error() string {

	if atomic.LoadInt32(&debug) == 0{

		return e.Info
	}
	frames := runtime.CallersFrames(e.stackPC)
	var (
		f      runtime.Frame
		more   bool
		result string
		index  int
	)
	for {
		f, more = frames.Next()
		if index = strings.Index(f.File, "src"); index != -1 {
			// trim GOPATH or GOROOT prifix
			f.File = string(f.File[index+4:])
		}
		result = fmt.Sprintf("%s%s\n\t%s:%d\n", result, f.Function, f.File, f.Line)
		if !more {
			break
		}
	}
	return fmt.Sprintf("%s\n%s",e.Info,result)
}

func New(s interface{})error{
	pcs := make([]uintptr, 32)
	count := runtime.Callers(2, pcs)
	return &Error{
		Info:fmt.Sprint(s),
		stackPC:pcs[:count],
		}
}

func Errorf(format string, a ...interface{}) error {
	pcs := make([]uintptr, 32)
	count := runtime.Callers(2, pcs)

	return &Error{

		Info:fmt.Sprintf(format, a...),
		stackPC:pcs[:count],

	}
}

func Errorln(a ...interface{}) error {
	pcs := make([]uintptr, 32)
	count := runtime.Callers(2, pcs)
	return &Error{
		Info:fmt.Sprintln(a...),
		stackPC:pcs[:count],

	}
}


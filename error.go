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

//判断两错误是否同源相等
func Equal(e1,e2 error)bool{
	if e1 == e2 {return true}
	if e1 == nil || e2 == nil{return false}
	E1,ok1 :=e1.(*Error)
	E2,ok2 :=e2.(*Error)
	if ok2 && !ok1{return E2.rawErr == e1}
	if ok1 && !ok2{return E1.rawErr == e2}

	if E1.rawErr == E2.rawErr && e1!=nil{
		return true
	}
	if E1.Info == E2.Info && E1 == E2 {
		return true
	}
	return false
}

type Errorer interface {
	Error() *Error
}

type Error struct {
	Info string
	rawErr error     // 保存原始错误信息
	stackPC []uintptr // 保存函数调用栈指针
}

func (e *Error) MarshalJSON() ([]byte, error){
	return []byte(fmt.Sprintf(`"%s"`,e.Info)),nil
}

func (e *Error) UnmarshalJSON(d []byte) error{
	e.Info = string(d)
	return nil
}

func (e *Error) IsNil() bool {
	return e.Info == ""
}
func (e *Error) String() string {
	return e.Info
}
//
func (e *Error) MarkPos() *Error {
	pcs := make([]uintptr, 32)
	count := runtime.Callers(2, pcs)
	return &Error{
		Info:e.Info,
		stackPC:pcs[:count],
		rawErr:e.rawErr,
	}
}

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

func New(s interface{})*Error{
	pcs := make([]uintptr, 32)
	count := runtime.Callers(2, pcs)
	e:=&Error{
		Info:fmt.Sprint(s),
		stackPC:pcs[:count],
	}
	switch raw:= s.(type){
	case error:
		e.rawErr = raw
	default:
		e.rawErr = e
	}
	return e
}


func Errorf(format string, a ...interface{}) *Error {
	pcs := make([]uintptr, 32)
	count := runtime.Callers(2, pcs)
	e:=&Error{
		Info:fmt.Sprintf(format, a...),
		stackPC:pcs[:count],
	}
	e.rawErr = e
	return e
}

func Errorln(a ...interface{}) *Error {
	pcs := make([]uintptr, 32)
	count := runtime.Callers(2, pcs)
	e := &Error{
		Info:fmt.Sprintln(a...),
		stackPC:pcs[:count],
	}
	e.rawErr = e
	return e
}


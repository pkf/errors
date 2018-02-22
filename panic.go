package errors

import (
	"bytes"
	"runtime"
)

type (
	PrintlnFn func(a ...interface{}) (n int, err error)
	PrintFn func(a ...interface{}) (n int, err error)
	PrintfFn func(format string, a ...interface{}) (n int, err error)
)

func RecoverPrintln(p PrintlnFn) {
	e := recover()
	if e != nil {
		p(string(readStack()))
	}
}
func RecoverPrint(p PrintFn) {
	e := recover()
	if e != nil {
		p(string(readStack()))
	}
}

//错误信息占最后一个%s
func RecoverPrintf(p PrintfFn, fmt string, a ...interface{}) {
	e := recover()
	if e != nil {
		p(fmt, append(a, string(readStack()))...)
	}
}

func RecoverToError(err *error) {
	e := recover()
	if e != nil {
		buf := readStack()
		*err = Err_PanicRecover.New(buf)
	}
}
func RecoverFn(fn func(err error)) {
	e := recover()
	if e != nil {
		buf := readStack()
		if fn != nil {
			fn(Err_PanicRecover.New(buf))
		}
	}
}
func Recover() {
	recover()
}

func readStack() []byte {
	buf := make([]byte, 1<<10)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			buf = buf[:n]
			break
		}
		buf = make([]byte, 2*len(buf))
	}

	i := bytes.Index(buf, []byte("/src/runtime/panic.go:"))
	if i > 0 {
		m := bytes.IndexByte(buf[i:], '\n')
		if m > 0 {
			buf = buf[i+m:]
		}
	}
	return buf
}

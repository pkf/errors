package errors

import (
	"fmt"
	"runtime"
)

type PreDefineCode int

var (
	Err_Normal       = PreDefineCode(-1) //通用错误
	Err_Timeout      = PreDefineCode(-2) //超时错误
	Err_Parameter    = PreDefineCode(-3) //参数错误
	Err_Option       = PreDefineCode(-4) //配置错误
	Err_PanicRecover = PreDefineCode(-5) //接受到panic错误

	Err_IPBlocked      = PreDefineCode(-21) //IP被封
	Err_UnexpectedPage = PreDefineCode(-22) //非预期页面
	Err_NetworkFault   = PreDefineCode(-23) //网络故障
	Err_ParseFailed    = PreDefineCode(-24) //解析失败
	Err_FareIsCache    = PreDefineCode(-25) //票价是缓存
	Err_NotFoundFlight = PreDefineCode(-26) //没有找到该航班

	Err_NotSupport     = PreDefineCode(-31) //不支持
	Err_InvalidPacket  = PreDefineCode(-32) //数据包错误
	Err_BufferOverflow = PreDefineCode(-33) //缓冲溢出

	Err_NotFound = PreDefineCode(-40) //没有找到该资源

)

func (e PreDefineCode) Int() int {
	return int(e)
}
func (e PreDefineCode) New(s interface{}) *Error {
	if s == nil{
		return nil
	}
	pcs := make([]uintptr, 32)
	count := runtime.Callers(2, pcs)
	if s1, ok := s.([]byte); ok {
		s = string(s1)
	}
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

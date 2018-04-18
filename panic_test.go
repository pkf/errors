package errors

import (
	"fmt"
	"testing"

	"git.jiatu365.com/go/console"
)

func TestRecover(t *testing.T) {
	defer Recover()
	panic("123")

}

func TestRecoverFn(t *testing.T) {
	var e error
	defer func() {
		fmt.Println("---------------")
		fmt.Println(e)
		fmt.Println("---------------")
	}()
	defer RecoverFn(func(err error) { e = err })
	panic("123")

}
func TestRecoverPrintln(t *testing.T) {
	defer RecoverPrintln(fmt.Println)
	panic("123")
}
func TestRecoverPrint(t *testing.T) {
	defer RecoverPrint(fmt.Print)
	panic("123")
}

func TestRecoverPrintf(t *testing.T) {
	defer RecoverPrintf(console.Printf, "%s:%s", console.HiRed("Panic"))
	panic("123")
}

func TestRecoverToError(t *testing.T) {
	var e error
	defer func() {
		fmt.Println("---------------")
		fmt.Println(e)
		fmt.Println("---------------")
	}()
	defer RecoverToError(&e)
	panic("123")

}

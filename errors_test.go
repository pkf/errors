package errors

import (
	"testing"
	"fmt"
)

func TestPrint(t *testing.T) {
	EnableDebug()
	err := printCase()
	fmt.Println(err)
}

var case1 = New("error case 1")
func printCase()error{
	return case1.MarkPos()
}
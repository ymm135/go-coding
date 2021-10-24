//testlib.go
package main
/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -ltest
#include <stdio.h>
#include <stdlib.h>
#include "test.h"
*/
import "C"

import "unsafe"
import "fmt"

func main() {
	fmt.Println("==c library test ==")
	fmt.Printf("rannum:%x\n", C.random())

	cs := C.CString("Hello world")
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
	//C.free(unsafe.Pointer(cs))
	C.fflush((*C.FILE)(C.stdout))

	fmt.Println("")
	fmt.Println(C.Value)
	C.sayHello()
}

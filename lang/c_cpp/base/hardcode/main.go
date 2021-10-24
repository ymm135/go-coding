package main

//#include <stdio.h>
//#include <stdlib.h>
/*
void Hello(char *str) {
    printf("%s\n", str);
}
*/
import "C"
import "unsafe"

/**
go调用clamav的api进行文件查杀
*/

func main() {
	s := "Hello Cgo"
	cs := C.CString(s)
	C.Hello(cs)
	C.free(unsafe.Pointer(cs))
}

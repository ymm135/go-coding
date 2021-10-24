# go调用C/C++ 
在go模块中使用c/c++代码、静态库、动态库, 可以把Go模块作为逻辑组织者,c/c++作为功能模块.

# CGO
要使用CGO特性，需要安装C/C++构建工具链，在macOS和Linux下是要安装GCC，在windows下是需要安装MinGW工具。同时需要保证环境变量CGO_ENABLED被设置为1，这表示CGO是被启用的状态。在本地构建时CGO_ENABLED默认是启用的，当交叉构建时CGO默认是禁止的。比如要交叉构建ARM环境运行的Go程序，需要手工设置好C/C++交叉构建的工具链，同时开启CGO_ENABLED环境变量。然后通过import "C"语句启用CGO特性。  

## import "C"语句
如果在Go代码中出现了import "C"语句则表示使用了CGO特性，紧跟在这行语句前面的注释是一种特殊语法，里面包含的是正常的C语言代码。当确保CGO启用的情况下，还可以在当前目录中包含C/C++对应的源文件。  
```
package cgo_helper

//#include <stdio.h>
import "C"

type CChar C.char

func (p *CChar) GoString() string {
    return C.GoString((*C.char)(p))
}

func PrintCString(cs *C.char) {
    C.puts(cs)
}
```

## #cgo语句
在import "C"语句前的注释中可以通过#cgo语句设置编译阶段和链接阶段的相关参数。编译阶段的参数主要用于定义相关宏和指定头文件检索路径。链接阶段的参数主要是指定库文件检索路径和要链接的库文件。  
```
// #cgo CFLAGS: -DPNG_DEBUG=1 -I./include
// #cgo LDFLAGS: -L/usr/local/lib -lpng
// #include <png.h>
import "C"
```

## CGO问题  
[参考链接](https://cloud.tencent.com/developer/article/1650525)  
go 语言提供了这样的工具，叫做 "cgo", 你也可以用 swig 之类的工具生成大量胶水代码，但是它的核心还是 cgo，但是很快你会发现，事情其实没那么简单 (不同于 lua 和 cpython 等使用 c 开发的解释语言)。最广泛流传的一篇警告来自 go 语言的作者之一 Dave Cheney, cgo is not Go, 这篇文章告诫我们，cgo 的缺点很多：

1. 编译变慢，实际会使用 c 语言编译工具，还要处理 c 语言的跨平台问题
1. 编译变得复杂
1. 不支持交叉编译
1. 其他很多 go 语言的工具不能使用
1. C 与 Go 语言之间的的互相调用繁琐，是会有性能开销的
1. C 语言是主导，这时候 go 变得不重要，其实和你用 python 调用 c 一样
1. 部署复杂，不再只是一个简单的二进制  

**具体还要了解cgo的实现方式及原理**  

# go模块中使用c/c++代码
> c代码需要用//或/* */注释，并且和import "C"之间不能有空行，必须紧接其后。

```
//cgotest.go
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
func main() {
    s := "Hello Cgo"
    cs := C.CString(s)
    C.Hello(cs)
    C.free(unsafe.Pointer(cs))
}
```



# go模块中使用c/c++静态库



# go模块中使用c/c++动态库
首先编译动态库
```
//test.h                                                                             
#ifndef __TEST_H 
#define __TEST_H 
#ifdef __cplusplus
extern "C" {
#endif
extern int Value; 
extern void sayHello();
extern int add(int a, int b);
#ifdef __cplusplus 
} 
#endif
#endif
```

```
//test.c
#include <stdio.h> 
int Value = 123456;
void sayHello()
{
    printf("this is a c library.");
}

int add(int a, int b)
{
    return a + b;
}
```

编译动态库`gcc -fPIC -shared -o libtest.so test.c`  

需要配置动态库的搜索路径`LD_LIBRARY_PATH`,也可以拷贝到系统库下  

通过CFLAGS配置编译选项(头文件)，通过LDFLAGS来链接指定目录下的动态库。注意import "C"是紧挨着注释的，没有空行。  

```
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
    
    fmt.Println("");
    fmt.Println(C.Value)
    C.sayHello()
}
```


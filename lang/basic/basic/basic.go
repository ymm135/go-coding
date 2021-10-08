package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 包内变量，没有全局变量的概念, 声明多个是可以使用括号代替
var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	// go语言都是有初始值的 // 变量定义必须要使用的
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s) // %s 不能打印字符串“”,只能打印空值
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	// 类型推断，C++后续也是支持的
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

// 推荐使用这种方式
func variableShorter() {
	// 简介的写法,使用: 代替var
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	// 可以写类型，也可以不写
	const name string = "xiaoming"
	fmt.Println(name)

	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int = -1
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()
	consts()
	enums()
}

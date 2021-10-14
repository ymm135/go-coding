package main

import "fmt"

// 问题就是sum是方法申请局部变量, 已经出现逃逸
func adder() func(int) int { 
	sum := 0
	return func(v int) int {
		sum += v           //自由变量
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder() //is trivial and also works. // dlv打印 &a = (*func(int) int)(0xc0000a5ed8)
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
		//fmt.Println(adder()(i)) //这样输出的就是0,1,2,3,4,5,6,7,8,9
	}

	a1 := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a1 = a1(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n",
			i, s)
	}
}

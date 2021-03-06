package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//多个返回值
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	// q = a / b
	// r = a % b
	return a / b, a % b //建议方式
}

// 函数作为参数
func apply(op func(int, int) int, a, b int) int {
	//获取函数的指针
	p := reflect.ValueOf(op).Pointer()
	//获取函数名
	opName := runtime.FuncForPC(p).Name()

	fmt.Printf("Calling function %s with args "+
		"(%d, %d)\n", opName, a, b)

	return op(a, b)
}

// 可变参数, 我的猜测是range可以操作len(type)的所有类型
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

// 也可以通过指针的方式返回 swap( a, b *int)
func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println("Error handling")
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	fmt.Printf("13 div 3 is %d mod %d\n", q, r)

	fmt.Println("pow(3, 4) is:", apply(
		func(a int, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))

	fmt.Println("1+2+...+5 =", sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println("a, b after swap is:", a, b)
}

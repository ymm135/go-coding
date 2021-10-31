package main

import (
	"fmt"
	r "imooc.com/ccmouse/learngo/lang/reflect"
	"reflect"
)

/**
学习Go反射,主要为了搭建框架,比如动态代理，生命周期管理，参数自动填充等
*/
func main() {

	// With reflection
	fT := reflect.TypeOf(r.Foo{})
	fmt.Println("numFiled:", fT.NumField(), "numMethod:",  fT.NumMethod())

	helloMethod, isExist := fT.MethodByName("Hello")
	if isExist {
		helloMethod.Func.Call([]reflect.Value{reflect.ValueOf("Direct")})
	}

	fV := reflect.New(fT)

	m := fV.MethodByName("Hello")
	if m.IsValid() {
		m.Call([]reflect.Value{reflect.ValueOf("New")})
	}

	// 动态代理

}

func testProxy()  {

}
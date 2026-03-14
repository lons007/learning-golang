package main

import (
	"fmt"
	"learning-golang/src/mypkg"
)

func main() {
	var value mypkg.MyInterface
	value = mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameters(127.3)
	fmt.Println(value.MethodWithReturnValue())
	fmt.Println(value)
}

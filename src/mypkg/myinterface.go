package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameters(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}

func (m MyType) MethodWithParameters(f float64) {
	fmt.Println("MethodWithParameters called with", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}

func (m MyType) MethodNotInterface() {
	fmt.Println("MethodNotInterface called")
}
package main

import (
	"fmt"

)

type Number int
type DoubleNumbers int

func (n *Number) Display() {
	fmt.Println(*n)
}

func (n *Number) Double() {
	*n *= 2
}

func (n *Number) DoubleNumbers() DoubleNumbers{
	return DoubleNumbers(*n * 2)
}

func main() {
	number := Number(4)
	number.Display()
	number.Double()
	number.Display()
	fmt.Println(number.DoubleNumbers())
	// number.DoubleNumbers()
	number.Display()

	// fmt.Println("Original value of number:", number)
	// fmt.Println("number after calling Double:", number)
}
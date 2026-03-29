package main

import "fmt"

type Counter struct {
	TestCount int
}

func (c Counter) GetValue() int {
	return c.TestCount
}

func (c *Counter) Increment() {
	c.TestCount++
}

func (c *Counter) Decrement() {
	if c.TestCount > 0 {
		c.TestCount--
	}
	
}

func (c *Counter) Reset() {
	c.TestCount = 0
}

func main() {
	var c Counter

	c.Increment()
	c.Increment()
	c.Increment()
	c.Decrement()
	fmt.Printf("Счетчик: %d\n",c.GetValue())

}

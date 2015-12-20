package main

import "fmt"

type Calculator interface {
	Calculate(a int, b int) int
}

type Add struct{}

func (x Add) Calculate(a int, b int) int {
	return a + b
}

type Sub struct{}

func (x Sub) Calculate(a int, b int) int {
	return a - b
}

func main() {
	var add Add
	var sub Sub
	var calc Calculator

	fmt.Println(add.Calculate(1, 3))
	calc = add
	fmt.Println(calc.Calculate(1, 3))

	fmt.Println(sub.Calculate(1, 3))
	calc = sub
	fmt.Println(calc.Calculate(1, 3))
}

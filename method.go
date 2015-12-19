package main

import "fmt"

type myType int

func main() {
	var z myType = 1234
	z.println()

	z.setByValue(100)
	z.println()

	z.setByPointer(100)
	z.println()
}

func (value myType) println() {
	fmt.Println(value)
}

// レシーバが、ポインタではないので値は変わらない
func (value myType) setByValue(setValue myType) {
	value = setValue
}

// レシーバが、ポインタなので値が変わる
func (value *myType) setByPointer(setValue myType) {
	*value = setValue
}

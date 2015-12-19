package main

import "fmt"

type myType int

func main() {
	var z myType

	fmt.Println(z.add(3))

	var mv = z.add
	fmt.Println(mv(3))
}

func (value *myType) add(i myType) myType {
	*value += i

	return *value
}

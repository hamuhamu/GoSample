package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{name: "はむ", age: 22}

	fmt.Println(p)
}

package main

import "fmt"

type embeded struct {
	i int
}

func (x embeded) do() {
	fmt.Println(x.i)
}

type test struct {
	embeded
}

func main() {
	var x test
	x.i = 10

	x.do()
	fmt.Println(x.i)
}

package main

import "fmt"

// 0, 1, 2, 3という値が入る
const (
	ZERO = iota
	ONE
	TWO
	THREE
)

// 0, 2, 4という値が入る
const (
	HOGE = iota * 2
	FUGA
	PIYO
)

func main() {
	fmt.Println(ZERO)
	fmt.Println(ONE)
	fmt.Println(TWO)
	fmt.Println(THREE)

	fmt.Println("----------------------------")

	fmt.Println(HOGE)
	fmt.Println(FUGA)
	fmt.Println(PIYO)
}

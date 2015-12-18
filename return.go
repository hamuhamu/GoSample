package main

import "fmt"

func main() {
	add, sub, mul, div := calc(3, 2)

	fmt.Println(add)
	fmt.Println(sub)
	fmt.Println(mul)
	fmt.Println(div)
}

// 戻り値に名前をつけれる
// returnを省略できる
func calc(a int, b int) (add int, sub int, mul int, div float32) {
	add = a + b
	sub = a - b
	mul = a * b
	div = float32(a) / float32(b)

	return
}

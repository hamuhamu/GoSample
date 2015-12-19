package main

import "fmt"

type MyData struct {
	s   string
	i   int
	int // 匿名フィールド
}

func main() {
	var x MyData
	x.s = "はむはむ"
	x.i = 10
	x.int = 100

	fmt.Println(x)
}

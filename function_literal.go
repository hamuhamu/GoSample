package main

import "fmt"

func main() {

	val := 123

	// 関数リテラル(匿名関数)
	func(i int) {
		// 関数リテラル外の変数にアクセスできる
		fmt.Println(i * val)
	}(10)

	f := func(s string) {
		fmt.Println(s)
	}

	f("hoge")
}

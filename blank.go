package main

import "fmt"

func main() {
	// この書き方ではbが使われていないため、コンパイルエラーとなる
	// Goは使用していない変数を宣言しているとコンパイルエラーとする仕様である
	// var a, b int = hoge()

	// 使用しない変数はblank識別子で受け取るとよい
	// ブランク識別子は、参照できない/dev/nullのようなもの
	var a, _ int = hoge()

	fmt.Println(a)
}

func hoge() (int, int) {
	return 1, 2
}

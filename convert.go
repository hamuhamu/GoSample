package main

import "fmt"

func main() {
	var i int = 100
	// intからstringへ
	var str string = string(i)

	// "100"という文字列ではなくその整数のユニコードポイントする文字列となる
	fmt.Println(str)
}

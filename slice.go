package main

import "fmt"

func main() {
	x := [5]string{"a", "b", "c", "d", "e"}

	// スライス型の変数を宣言
	var s1 []string

	s1 = x[:]
	fmt.Println(s1)

	s2 := x[1:4]
	fmt.Println(s2)

	s3 := x[3:]
	fmt.Println(s3)

	s4 := x[:4]
	fmt.Println(s4)

	// スライスは、参照渡し(アドレス)の為、もとの配列の中身も変えてしまう
	s4[0] = "x"
	fmt.Println(x)
}

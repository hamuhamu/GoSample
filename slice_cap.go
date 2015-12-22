package main

import "fmt"

func main() {
	values := [...]int{0, 1, 2, 3, 4}
	s1 := values[1:3]

	fmt.Println("s1:  ", s1)
	fmt.Println("len: ", len(s1))
	fmt.Println("cap: ", cap(s1))

	// スライスをスライス
	s2 := s1[1:4]
	fmt.Println("s2:  ", s2)
	fmt.Println("len: ", len(s2))
	fmt.Println("cap: ", cap(s2))

	/* スライスからスライスを作る際は、キャパシティの値までスライスできる */

	// スライスの第３引数は、キャパシティ
	// キャパシティを指定してスライス
	s3 := s1[1:3:4]
	fmt.Println("s3:  ", s3)
	fmt.Println("len: ", len(s3))
	fmt.Println("cap: ", cap(s3))
}

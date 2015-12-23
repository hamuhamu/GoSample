package main

import "fmt"

func main() {
	dst := []int{1, 2, 3, 4}
	src := []int{5, 6, 7}

	// 戻り値(コピーされた要素数) = copy(コピー先スライス, コピー元スライス)
	count := copy(dst[2:], src)

	fmt.Println("dst: ", dst)
	fmt.Println("count: ", count)
}

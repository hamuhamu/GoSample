package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}
	test(s...)
	test("a", "b", "c")
}

// スライスを関数に渡す
func test(s ...string) {
	fmt.Println(len(s), s)
}

package main

import "fmt"

// main関数より先に呼ばれる
func init() {
	fmt.Println("init関数だよ")
}

func main() {
	fmt.Println("main関数だよ")
}

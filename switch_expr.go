package main

import "fmt"

func main() {

	for i := -2; i < 2; i++ {

		// switchのcaseに式を使用する
		switch true {
		case i < 0:
			fmt.Println("-")
		case i > 0:
			fmt.Println("+")
		default:
			fmt.Println(i)
		}
	}
}

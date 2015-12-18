package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		switch i {
		case 1:
			fmt.Println("1")
		case 2, 3:
			fmt.Println("2か３")
		default:
			fmt.Println("その他")
		}
	}
}

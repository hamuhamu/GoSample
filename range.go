package main

import "fmt"

func main() {
	ary := [...]int{1, 2, 3, 4, 5}

	for i := range ary {
		fmt.Println(i)
	}
}

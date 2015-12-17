package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println("こんにちは", i, "回目")
	}

	fmt.Println("--------------------------")

	i := 1
	for i <= 5 {
		fmt.Println("こんにちは", i, "回目")
		i++
	}

	fmt.Println("--------------------------")

	i = 1
	for {
		// breakしないと無限ループする
		if i > 5 {
			break
		}

		fmt.Println("こんにちは", i, "回目")
		i++
	}
}

package main

import "fmt"

func main() {
	variadic(1, "こんにちわ")
	variadic(2, "Ruby", "Go")
	variadic(3, "very", "general", "specific")
}

// wordsは、可変長引数
func variadic(number int, words ...string) {
	fmt.Println("number:", number)

	for _, word := range words {
		fmt.Println("word:", word)
	}
}

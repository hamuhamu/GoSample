package main

import "fmt"
import "unicode/utf8"

func main() {
	var ja string = "Go言語だよ"

	fmt.Println(ja, "len(ja)", len(ja))
	// UTF8文字列をカウントするにはutf8.RuneCountInStringを使う
	fmt.Println(ja, "utf8.RuneCountInString(ja)", utf8.RuneCountInString(ja))
}

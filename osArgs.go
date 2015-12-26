package main

import (
	"fmt"
	"os"
)

// Usage:
// go run osArgs.go hoge fuga
// hoge fuga
func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

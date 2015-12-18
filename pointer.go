package main

import "fmt"

func main() {
	// ポインタの宣言
	var ptr *int
	i := 12345

	ptr = &i

	fmt.Println("iのアドレス", &i)
	fmt.Println("ptrの値", ptr)

	fmt.Println("iの値", i)
	fmt.Println("ポインタ経由のiの値", *ptr)

	*ptr = 999

	fmt.Println("ポインタ経由ので変更したiの値", *ptr)
}

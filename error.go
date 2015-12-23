package main

import "fmt"

type MyError struct {
	message string
}

func (err MyError) Error() string {
	return err.message
}

func main() {
	val, err := hoge(1)
	fmt.Println(val, err)

	val, err = hoge(101)
	fmt.Println(val, err)
}

func hoge(v int) (result bool, err error) {
	if v > 100 {
		return false, MyError{"値が大きすぎます"}
	}
	result = true

	return
}

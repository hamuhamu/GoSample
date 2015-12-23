package main

import "fmt"

func main() {
	f2()
}

func f1() {
	panic("パニック発生!!")
}

func f2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("パニック中断")
		}
		fmt.Println("ここは、実行される")
	}()

	f1()
	fmt.Println("ここは、実行されない")
}

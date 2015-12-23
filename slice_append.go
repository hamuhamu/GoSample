package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := append(s1, 5, 6)
	// スライスにスライスを加える際は、...が必要
	s3 := append(s2, s1...)

	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	fmt.Println("s3: ", s3)
}

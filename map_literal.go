package main

import "fmt"

func main() {
	// map初期化
	hoge := map[string]string{}
	fmt.Println(hoge)

	capitals := map[string]string{
		"日本":   "東京",
		"アメリカ": "ワシントンD.C.",
		"中国":   "北京",
	}

	fmt.Println(capitals)
}

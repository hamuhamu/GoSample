package main

import "fmt"

func main() {
	// mapを作成
	capitals := make(map[string]string)

	capitals["日本"] = "東京"
	capitals["アメリカ"] = "ワシントンD.C."
	capitals["中国"] = "北京"

	fmt.Println(capitals["日本"])
	fmt.Println(capitals["アメリカ"])
	fmt.Println(capitals["中国"])

	fmt.Println()

	for country, capital := range capitals {
		fmt.Println(country, capital)
	}

	// マップのキー存在確認
	_, ok := capitals["イギリス"]

	if ok {
		fmt.Println("キーが存在する")
	} else {
		fmt.Println("キーが存在しない")
	}
}

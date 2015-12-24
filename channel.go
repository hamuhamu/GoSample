package main

import "fmt"

func main() {
	// チャネルの作成
	c := make(chan int)

	// 送信専用チャネル
	go func(s chan<- int) {

		// チャネルへ0~4の値を順番に送信
		for i := 0; i < 5; i++ {
			s <- i
		}

		close(s)
	}(c)

	for {
		// チャネルから受信
		// チャネルがクローズされるとokにfalseが入る
		value, ok := <-c

		if !ok {
			break
		}

		fmt.Println(value)
	}
}

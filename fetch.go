package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// HTMLのbodyを取得する
func main() {
	url := "http://gopl.io"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s", b)
}

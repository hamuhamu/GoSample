package main

func main() {
	f()
}

// panicが発生すると、関数呼び出しを順次遡りプログラムを終了させる
func f() {
	panic("パニック発生!!")
}

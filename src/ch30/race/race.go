package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan string, 3) // 通过 buffer channel 防止协程阻塞
	source := [...]string{"baidu", "bing", "google"}
	res := ""
	for _, value := range source {
		go func(source string) {
			fetchMessage(ch, source)
		}(value)
	}

	for i := 0; i < len(source); i++ {
		res += <-ch + "\n"
	}
	fmt.Println("before::", runtime.NumGoroutine())
	fmt.Println(res)
	fmt.Println("After::", runtime.NumGoroutine())
}

func fetchMessage(ch chan string, source string) {
	time.Sleep(time.Microsecond * 100)
	fmt.Println(source)
	ch <- fmt.Sprintf("get message from %s", source)
}

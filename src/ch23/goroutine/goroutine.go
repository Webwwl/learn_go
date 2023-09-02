package main

import (
	"fmt"
	"time"
)

func main() {
	GoRoutingDemo()
}

func GoRoutingDemo() {
	for i := 1; i <= 10; i++ {
		go func(i int) {
			fmt.Println("current value", i)
		}(i)
	}
	time.Sleep(time.Microsecond * 300) // sleep 一段时间，不然 main 退出了终端看不到协程的输出
}

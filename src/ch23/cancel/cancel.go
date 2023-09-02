package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{}, 0)
	receive(ch)
	// cancel(ch)  // 不用 close 的话，有多少 receive 就需要调用多少次 cancel(ch) 耦合较重
	// cancel(ch)
	// cancel(ch)
	// cancel(ch)
	// cancel(ch)
	cancel2(ch)
	time.Sleep(time.Second * 1)
}

func receive(ch chan struct{}) {
	for i := 0; i < 5; i++ {
		go func(i int, ch chan struct{}) {
			for {
				if isCanceled(ch) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println("canceled:", i)
		}(i, ch)
	}
}

func isCanceled(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func cancel(ch chan struct{}) {
	ch <- struct{}{}
}

func cancel2(ch chan struct{}) {
	close(ch)
}

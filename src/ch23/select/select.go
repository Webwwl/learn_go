package main

import (
	"fmt"
	"time"
)

func main() {
	SelectDemo()
}

func SelectDemo() {
	select {
	case ret := <-AsyncTask():
		fmt.Println("get data:", ret)
	case <-time.After(time.Millisecond * 100):
		panic("timeout")
	}
	time.Sleep(time.Second * 1)
}

func AsyncTask() chan string {
	ret := make(chan string)
	go func() {
		time.Sleep(time.Millisecond * 200)
		ret <- "hello"
	}()
	return ret
}

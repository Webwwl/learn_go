package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ret := Task1()
	Task2()
	fmt.Println(<-ret)
	wg.Wait()
}

func Task1() chan string {
	var wg sync.WaitGroup
	wg.Add(1)
	ret := make(chan string)
	go func() {
		fmt.Println("task1 start")
		time.Sleep(time.Microsecond * 200)
		ret <- "hello"
		fmt.Println("task2 end")
	}()
	return ret
}

func Task2() {
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("task2 start")
	time.Sleep(time.Microsecond * 100)
	fmt.Println("task2 end")
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	Producer(ch, &wg)
	Receive(ch, &wg)
	Receive(ch, &wg)
	wg.Wait()
}

func Producer(ch chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func Receive(ch chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			if value, ok := <-ch; ok {
				fmt.Println("get value:", value)
			} else {
				fmt.Println("get value end")
				break
			}
		}
		wg.Done()
	}()
}

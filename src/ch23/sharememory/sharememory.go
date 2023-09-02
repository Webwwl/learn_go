package main

import (
	"fmt"
	"sync"
)

func main() {
	ShareMemoThreadSafe()
}

func ShareMemoThreadSafe() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	count := 0
	for i := 1; i <= 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mutex.Unlock()
			}()
			mutex.Lock()
			count++
			wg.Done()
		}()
	}
	wg.Wait() // 通过 WaitGroup 就不需要 sleep 了
	fmt.Println("result:", count)
}

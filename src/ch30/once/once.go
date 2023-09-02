package main

import (
	"fmt"
	"sync"
	"unsafe"
)

type Obj struct {
}

var once sync.Once
var obj Obj

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ob := getObject()
			fmt.Printf("%x \n", unsafe.Pointer(&ob))
			wg.Done()
		}()
	}
	wg.Wait()
}

func getObject() Obj {
	once.Do(func() {
		fmt.Println("create obj")
		obj = *new(Obj)
	})
	return obj
}

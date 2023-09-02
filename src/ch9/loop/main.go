package main

import (
	"fmt"
)

func main() {
	Loop()
}

func Loop() {
	for n := 0; n < 5; n++ {
		fmt.Println(n)
	}
}

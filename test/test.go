package main

import (
	"fmt"
	"reflect"
)

func main() {
	testMapKeyOrder()
	testDeepEqual()
}

func testMapKeyOrder() {
	message := map[string]int{"a": 1, "b": 2, "c": 3}
	for i := 0; i < 10; i++ {
		keys := []string{}
		for k, _ := range message {
			keys = append(keys, k)
		}
		fmt.Println(keys)
	}
}

func testDeepEqual() {
	a1 := [...]string{"a", "b", "c"}
	a2 := [...]string{"a", "b", "c"}
	fmt.Println("a === b::", reflect.DeepEqual(a1, a2))
}

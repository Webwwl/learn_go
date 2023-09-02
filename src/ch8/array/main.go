package main

import "fmt"

func main() {
	ArrayInit()
	ArrayLoop()
	ArrayLoop2()
	ArraySection()
}

func ArrayInit() {
	var arr [3]int
	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3}
	arr3 := [3][1]int{{1}, {2}, {3}}
	arr[0] = 2
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

func ArrayLoop() {
	arr := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func ArrayLoop2() {
	arr := [...]int{1, 2, 3, 4, 5}
	for _, value := range arr {
		fmt.Println(value)
	}
}

func ArraySection() {
	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr[:])   // 所有
	fmt.Println(arr[0:4]) // 0 - 4 (不包括 4)
	fmt.Println(arr[1:])  // 1 之后的(包括 1)
	fmt.Println(arr[:3])  // 0 - 3 (不包括 3)
}

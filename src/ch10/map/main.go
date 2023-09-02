package main

import "fmt"

func main() {
	MapInit()
	MapInitKeyNotExist()
	MapLoop()
	MapWithFuncValue()
}

func MapInit() {
	m := map[string]int{}
	m["0"] = 1
	fmt.Println(m)
	m1 := map[string]int{"1": 1, "2": 2}
	fmt.Println(m1)
}

func MapInitKeyNotExist() {
	m := map[string]int{}
	fmt.Println(m["2"]) // key不存在的时候，打印的 value==0(针对申明的 value 是 int 时)
	m["2"] = 1
	if v, ok := m["2"]; ok { // 当一个key 的值是 0 的时候，不确定是 key 不存在还是本来值就是 0，可以通过这种方式确定
		fmt.Println("key 2 value is ", v)
	} else {
		fmt.Println("key 2 not exist")
	}
}

func MapLoop() {
	m := map[string]int{"1": 1, "2": 2, "3": 3}
	for k, v := range m {
		fmt.Println("k v::", k, v)
	}
}

func MapWithFuncValue() {
	m := map[int]func(v int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	fmt.Println(m[1](2), m[2](2), m[3](2))
}

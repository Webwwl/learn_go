package main

import "fmt"

func main() {
	TestType(1)
	TestType("1")
	TestType(true)
}

func TestType(p interface{}) { // 空 interface 有类似断言的能力
	// if _, ok := p.(int); ok {
	// 	fmt.Println("get int", p)
	// 	return
	// }
	// if _, ok := p.(string); ok {
	// 	fmt.Println("get string", p)
	// 	return
	// }
	// fmt.Println("unknown type")
	switch p.(type) {
	case int:
		fmt.Println("get int", p)
	case string:
		fmt.Println("get string", p)
	default:
		fmt.Println("unknown type")
	}
}

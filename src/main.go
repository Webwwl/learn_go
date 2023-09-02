package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("get args:::", os.Args[1])
	} else {
		fmt.Println("Hello World")
	}
	// TestType()
	TestOperate()
	os.Exit(0)
}

type MyInt int64

func TestType() {
	var a int = 1
	var b int64
	b = int64(a)
	var c MyInt
	// c = b
	fmt.Println(a, b, c)
	d := 1
	e := &a
	fmt.Println(d, e)
	// e := e + 1 // 不允许内存地址运算
	var m string
	fmt.Println("*" + m + "*") // 字符串类型的初始值是空字符 不是 nil
}

func TestOperate() {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 5}
	fmt.Println(a == b)
	fmt.Println(a == d)
	// fmt.Println(a == c) // 长度不一样的数据不允许比较
}

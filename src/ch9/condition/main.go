package main

import (
	"fmt"
)

func main() {
	IfElse()
	Switch()
	Switch2()
}

func IfElse() {
	if a := 3; a > 2 { // 支持写表达式
		fmt.Println("a > 2")
	}
}

func Switch() {
	n := 0
	switch n {
	case 0, 1: // 支持写多个值
		fmt.Println("0011") // 默认会 break
	case 2:
		fmt.Println("2222")
	}
}

func Switch2() {
	price := 24
	switch { // 类似 if elseif else 的效果
	case price <= 10:
		fmt.Println(price)
	case price > 10 && price <= 20:
		fmt.Println(price - 2)
	case price > 20:
		fmt.Println(price - 5)
	}
}

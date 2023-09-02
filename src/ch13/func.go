package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a, b := getMultiValues()
	fmt.Println(a, b)
	fn := LogSpendTime(log)
	fn(10)
	sum1 := Sum(1, 2, 3, 4)
	sum2 := Sum(1, 2, 3, 4, 5)
	fmt.Println("sum result", sum1, sum2)
	TestDeffer()
}

func getMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func LogSpendTime(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start_time := time.Now()
		ret := inner(op)
		spend_time := time.Since(start_time).Seconds()
		fmt.Printf("func spend time %f second \n", spend_time)
		return ret
	}
}

func log(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func Sum(args ...int) int {
	ret := 0
	for _, value := range args {
		ret += value
	}
	return ret
}

func TestDeffer() {
	defer clear() // 有错误依然会执行
	fmt.Println("start")
	panic("throw some error") // 抛出错误
}

func clear() {
	fmt.Println("clear resource")
}

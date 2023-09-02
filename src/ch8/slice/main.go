package main

import "fmt"

func main() {
	Slice()
	SliceAutoExpand()
	SliceShareMemory()
}

func Slice() {
	s0 := []int{1, 2, 3}
	fmt.Println(s0) // cap(s0) 获取 s0 的容量
	s0 = append(s0, 4)
	fmt.Println(s0)
	s1 := make([]int, 3, 5)
	fmt.Println(len(s1))
	// fmt.Println(s1[3]) // 超出元素个数，报错
	s1 = append(s1, 4)
	fmt.Println(len(s1))
	fmt.Println(s1[3])
}

func SliceAutoExpand() {
	var s []int
	for i := 0; i < 10; i++ {
		fmt.Println(len(s), cap(s)) // 容量按 2 的指数递增
		s = append(s, i)            // 每次扩容后会指向新的地址，所以需要 s= 赋值
	}
}

func SliceShareMemory() {
	s := []string{"Mike", "Cindy", "Blues"}
	s1 := s[0:2]
	s2 := s[1:]
	fmt.Println(s1, s2)
	s1[1] = "Unknown"
	fmt.Println(s1, s2)
}

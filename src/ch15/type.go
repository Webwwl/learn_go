package main

import "fmt"

func main() {
	NewTest()
	TypeArgs()
}

type Student struct {
	id   int
	name string
	age  int
}

type A struct {
	a string
}

func (s Student) String() {
	fmt.Printf("name: %s age: %d\n", s.name, s.age)
}

func (s *Student) String2() {
	fmt.Printf("name: %s age: %d\n", s.name, s.age)
}

func NewTest() {
	stu := Student{name: "wwl", age: 10}
	stu2 := Student{1, "wwl", 10}
	fmt.Println(stu)
	stu.name = "Mike"
	fmt.Println(stu)
	fmt.Println(stu2)
}

func TypeArgs() {
	s := Student{1, "wwl", 10}
	LogStudentMessage((s))
	LogStudentMessage2(&s)
	s.String()
	s.String2()
}

func LogStudentMessage(s Student) {
	fmt.Printf("name: %s age: %d\n", s.name, s.age)
}

func LogStudentMessage2(s *Student) {
	fmt.Printf("name: %s age: %d \n", s.name, s.age)
}

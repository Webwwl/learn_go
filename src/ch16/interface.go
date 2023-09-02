package main

import "fmt"

func main() {
	d := Dog{}
	b := Bird{}
	logName(&d)
	logName(&b)
}

type Animal interface {
	getName() string
}

type Dog struct{}

func (d *Dog) getName() string {
	return "name is: Dog"
}

type Bird struct{}

func (d *Bird) getName() string {
	return "name is: Bird"
}

func logName(animal Animal) { // 面向接口编程
	fmt.Println(animal.getName())
}

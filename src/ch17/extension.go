package main

import "fmt"

func main() {
	d := new(Dog)
	d.Speak2()
}

type Animal struct {
}

func (a *Animal) Speak() {
	fmt.Println("...")
}

func (a *Animal) Speak2() {
	a.Speak()
	fmt.Println("speak 2")
}

type Dog struct {
	Animal
}

func (d *Dog) Speak() {
	fmt.Println("wang wang wang")
}

func (d *Dog) Speak2() {
	fmt.Println("wang wang wang")
}

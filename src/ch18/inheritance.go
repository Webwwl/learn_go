package main

import "fmt"

func main() {
	goProgrammer := new(GoProgrammer)
	jsProgrammer := new(JavaScriptProgrammer)
	PrintHelloWorldCode(goProgrammer)
	PrintHelloWorldCode(jsProgrammer)
}

type Programmer interface {
	HelloWorldCode() string
}

type GoProgrammer struct{}

func (p *GoProgrammer) HelloWorldCode() string {
	return "fmt.print(\"hello world\")"
}

type JavaScriptProgrammer struct{}

func (p *JavaScriptProgrammer) HelloWorldCode() string {
	return "console.log(\"hello world\")"
}

func PrintHelloWorldCode(p Programmer) {
	fmt.Println(p.HelloWorldCode())
}

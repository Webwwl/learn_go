package main

import (
	"fmt"
	"strings"
)

func main() {
	StringsMethods()
}

func StringsMethods() {
	s := "hello"
	parts := strings.Split(s, "")
	fmt.Println(parts)
	fmt.Println(strings.Join(parts, "-"))
	fmt.Println(strings.Repeat(s, 2))
}

package main

import (
	"fmt"

	"example.com/pkg/magic"
)

func main() {
	foo := magic.Foo{}
	res := foo.Process(5)

	fmt.Println("demo - tools.go")
	fmt.Printf("foo.Process = %d\n", res)
}

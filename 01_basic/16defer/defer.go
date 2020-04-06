package main

import (
	"fmt"
)

func main() {

	demo()
	defer foo()
	bar()
}
func demo() {
	fmt.Println("this is demo")
}
func foo() {
	fmt.Println("this is foo")
}
func bar() {
	fmt.Println("this is bar")
}

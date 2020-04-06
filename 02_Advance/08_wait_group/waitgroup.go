package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	go bar()
	go demo()
	go foo()
	wg.Add(3) // how much go_routines to be waited
	fmt.Println("hello world")
	wg.Wait() // it will wait
}
func demo() {

	fmt.Println("hello demo")
	wg.Done() // it will pop one added go routine from add()
}
func bar() {

	fmt.Println("hello bar")
	wg.Done() // it will pop one added go routine from add()
}
func foo() {

	fmt.Println("hello foo")
	wg.Done() // it will pop one added go routine from add()
}

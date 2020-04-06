package main

import (
	"fmt"
)

func main() {
	i := 1
	if i == 1 {
		fmt.Println("answer is true")
	} else {
		fmt.Println("answer is false")
	}
	switch i {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("answer is to be known")
	}
}

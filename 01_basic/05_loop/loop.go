package main

import "fmt"

func main() {
	i := 0
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}
	for {
		//	i := 0
		if i > 2 {
			break
		}
		fmt.Println(i)
		i++
	}
	for i < 2 {
		fmt.Println(i)
	}
}

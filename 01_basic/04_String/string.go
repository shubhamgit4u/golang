package main

import "fmt"

func main() {
	str := "shubham"
	str2 := "sharma"
	fmt.Print(str + str2)
	fmt.Printf("%d", len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
}

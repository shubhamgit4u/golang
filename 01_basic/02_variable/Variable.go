package main

import "fmt"

var b = 7
var z = 10

func main() {
	i := 2
	j := "GO"
	fmt.Println(z) 
	z := 11
	fmt.Println(z) 
	k := true
	{
		a := 8
		fmt.Println(a)
	}
	fmt.Print("i =", i, "\n")
	fmt.Print("j =", j, "\n")
	fmt.Print("k =", k, "\n")

	{
		b := 1
		fmt.Println(b)
	}
}

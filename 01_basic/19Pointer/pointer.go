package main

import (
	"fmt"
)

func main() {

	var i = 5
	var p *int
	p = &i
	fmt.Println(i)
	fmt.Println(*p)
	fmt.Println(p)

}

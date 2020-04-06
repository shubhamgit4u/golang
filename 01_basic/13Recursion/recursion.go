package main

import (
	"fmt"
)

func main() {

	i := demo(5)
	fmt.Println(i)

}
func demo(i int) bool {
	if i == 0 {
		return true
	}
	fmt.Println(i)
	i--
	return demo(i)
}

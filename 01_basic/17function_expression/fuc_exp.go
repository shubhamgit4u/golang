package main

import (
	"fmt"
)

func main() {
	fun := func() {
		var i = 9
		i++

		fmt.Println(i)
	}
	fun()
	fun()
}

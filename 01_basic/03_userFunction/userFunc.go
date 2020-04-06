package main

import "fmt"

func main() {
	num1 := 2
	num2 := 2
	num3, flag := divide(num1, num2)
	fmt.Println("result is ", num3, " with status ", flag)
}
func divide(num1, num2 int) (int, bool) {
	if num1 > 0 {
		return num1 / num2, true
	}
	return 0, false
}

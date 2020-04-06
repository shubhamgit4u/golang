package main

import "fmt"

type user struct {
	name  string
	age   int
	marks []int
}

func main() {
	akash := user{name: "shubham", age: 24}
	akash.marks = make([]int, 100)
	akash.marks[0] = 22
	akash.age = 21
	fmt.Println(akash.age)
	rohit := user{"rohit", 21, make([]int, 100)}
	fmt.Println(rohit.age)
	if rohit.age == akash.age {
		fmt.Println("you have same age")
	}
}

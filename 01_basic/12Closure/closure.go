package main

import "fmt"

type person struct {
	name  string
	age   int
	marks int
}

func main() {
	x := record()
	y := record()
	fmt.Println(x())
	fmt.Println(y())
}

func record() func() person {
	user := person{}
	fmt.Scanf("%s", &user.name)
	fmt.Scanf("%d", &user.age)
	fmt.Scanf("%d", &user.marks)

	return func() person {
		grace := 5
		user.marks += grace
		return user

	}
}

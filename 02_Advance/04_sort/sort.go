package main

import (
	"fmt"
	"sort"
)

func main() {
	ls := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	bo := sort.IntsAreSorted(ls)
	fmt.Println(bo)
	sort.Ints(ls)
	bol := sort.IntsAreSorted(ls)
	fmt.Println(bol)
	fmt.Println(ls)

}

//ok??

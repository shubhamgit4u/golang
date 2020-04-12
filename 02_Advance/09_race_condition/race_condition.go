package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mt sync.Mutex

func main() {
	counter := 0
	//fmt.Printf("go routine =%d and counter value is %d \n", runtime.NumGoroutine(), counter)
	n := 200
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			//mt.Lock()
			counter++
			//mt.Unlock

			wg.Done()
		}()
		fmt.Printf("go routine =%d and counter value is %d \n", runtime.NumGoroutine(), counter)
	}

	fmt.Println("end value", counter)
	wg.Wait()
}

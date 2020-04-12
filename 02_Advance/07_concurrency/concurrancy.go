//ok so what is concurrancy?
//it is most important core part of go bcz of it go is popular
//in 2006 intel lauch its first dual core processor , in 2007 gooogle develop goland
//so there basic idea is to work on core level
//other language like python,java developed before 2006 , so they not able to work on core level
// go is first language to work on core level
// ok so before golang what was scienerio ?
//means how language deal with cpu?
// when one instruction pass it consumes time like -> when i=i+1 in backend what happened?
// it load i from memory so there is time gap between loading that is wasted
// so they want we will send next instruction till i loads?
// so concept of parallelism comes , so parallelism means processing two instrction at a same time
// but wait there is one cpu how can you process two instrction at a same time?
// so before 2006 there is single core so they manges time gap between that loading by use of thread
// so all language except go works on this way
// after 2006 core develops then instruction passes to two different part
// now two instruction passes  to two cores simuntaniouslly but what happen when in first core
//int i passes and in second core i = i+1 passes , now in this case second core ideal for some time right?
// means now developer wants more work from cores
//yes in comparison of before 2006 it works faster but core still remain vacant while loading
// same problem was with single cpu now same problem with cores right?
// so what happens concurrancy comes which deals core of cpu
// so dealing with instruction is called concurrancy and doing work is called parallesim?
// in some cases both are same thing in some part different
// now what happen when in core 2 i =i+1 passes , when i loads it passes print i // now that is problem?
// and it is called race condition
// one more thing handling these interleaving not in developer side it depends on os means
// what passes at what time depend on os? so many time it will not give you error but gives wrong output
// so while making program it is human intelligence what to make concurrant and if you want in some case and problem
//happens then serailization can be done i.e only print i if i = i+1 finish
// so in golang we call these threads go routine
// when you run program like hello world a single main gorutine generate
// so next further go routine in developer side
//means by default these is always one go routine means one task
// one most important thing -> if main go routine finish its work while other go routine are processing
// it will comes out of processing
// so we can use time.sleep in main go routine //use wait_group
// wait group is syncronisation//
package main

import (
	"fmt"
	"runtime"
)

func main() {
	go bar()
	go demo()
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
	fmt.Println("hello world")
}
func demo() {

	fmt.Println("hello demo")
}
func bar() {

	fmt.Println("hello bar")

}

// int i ->interleave 1 or you can say subtask or thread or module anything
//i=i+1  ->interleeave 2 // one main go routine work in serialized way
//print i ->interleave 3

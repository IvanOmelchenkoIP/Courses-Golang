//
// What the Race Condition is and how they occur
//
// In concurrent programming, Race Condition is a condition when the outcome of certain tasks can't fully be predicted 
// because it depends on interleaving of these tasks, which is non-deterministic. 
// Race conditions only occur between the concurrent task, if there is some communication between them, like sharing common data. 
// This is because the interleavings are not deterministic, so we cannot predict the order in which the tasks will run 
// (different interleavings often have to be considered), we even cannot predict if one task will finish before the other one starts, 
// so we cannot know how the tasks really will communicate. In Golang, Go Runtime Scheduler manages interleaving of the goroutines 
// which are most similar to the threads in, typically, a single process.
// 
// In the example of the written program, there are two goroutines:
// - goroutineInc increments x and prints the result. 
// - goroutineSquare counts a square of current x and prints it and the result
//
// In this case, goroutineSquare depends on the value of x, which is changed by goroutineInc, 
// so the goroutines communicate and have a common variable that they use. 
// The results of work of the two goroutines mostly arrive inconsistently regarding one another. 
// The output also changes each time the program is run, which means there is different interleaving. 
// During testing output, I have also stumbled upon the situation, where the square is counted for the previous value of x, 
// which probably means that goroutineSquare was allowed to count a square of x, 
// but then Scheduler interrupted it for the full completion of a single cycle of goroutineInc (with incrementing x and printing it) 
// and only after that the goroutineSquare was allowed to print its value.
//


package main

import (
	"fmt"
	"time"
)

func goroutineInc() {
	for i := 0; i < 5; i++ {
		x++
		fmt.Printf("Current x value: %d\n", x)
		time.Sleep(1 * time.Millisecond)
		
	}
}

func goroutineSquare() {
	for i := 0; i < 5; i++ {
		y := x * x
		fmt.Printf("Current x value: %d; square of the current x value: %d\n", x, y)
		time.Sleep(1 * time.Millisecond)
	}

}

var x int 

func main() {
	go goroutineInc()
	go goroutineSquare()

	time.Sleep(100 * time.Millisecond)
}
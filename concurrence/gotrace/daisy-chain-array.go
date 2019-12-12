package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func f(left chan<- int, right <-chan int) {
	left <- 1 + <-right
}

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	const n = 100

	// first we construct an array of n+1 channels each being a 'chan int'
	var channels [n + 1]chan int
	for i := range channels {
		channels[i] = make(chan int)
	}

	// now we wire n goroutines in a chain
	for i := 0; i < n; i++ {
		time.Sleep(time.Microsecond * 10)
		go f(channels[i], channels[i+1])
	}

	// insert a value into the right-hand end
	time.Sleep(time.Microsecond * 10)
	go func(c chan<- int) { c <- 1 }(channels[n])

	// pick up the value emerging from the left-hand end
	fmt.Println(<-channels[0])
}

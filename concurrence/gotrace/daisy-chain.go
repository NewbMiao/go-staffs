package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	const n = 100
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		time.Sleep(time.Microsecond * 10)
		go f(left, right)
		left = right
	}
	time.Sleep(time.Microsecond * 10)
	go func(c chan int) {
		c <- 1
	}(right)
	fmt.Println(<-leftmost)
}

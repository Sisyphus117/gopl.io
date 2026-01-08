package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	const limit = 100_000_000
	for stage := 100; stage <= limit; stage *= 10 {
		runtime.GC()
		prev := make(chan int, 1)
		prev <- 1
		start := time.Now()
		for i := 0; i < stage; i++ {
			c := make(chan int)
			go func(prev <-chan int, next chan<- int) {
				cur := <-prev
				cur++
				next <- cur
			}(prev, c)
			prev = c
		}
		<-prev
		fmt.Printf("stage:%d ,time:%v\n", stage, time.Since(start))
	}

}

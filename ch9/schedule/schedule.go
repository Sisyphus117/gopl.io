package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	ticker := time.NewTicker(5 * time.Second)
	done := make(chan struct{})
	go func() {
		<-ticker.C
		fmt.Println("count:", count)
		done <- struct{}{}
	}()

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			msg := <-ch1
			count++
			ch2 <- msg
		}
	}()
	go func() {
		for {
			msg := <-ch2
			ch1 <- msg
		}
	}()

	ch2 <- "ping"
	<-done
}

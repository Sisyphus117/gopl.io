// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

// !+
func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	ticker := time.NewTicker(time.Second)
	count := 0
	text := make(chan string)
	done := make(chan struct{})
	go func() {
		defer close(text)
		for {
			select {
			case <-done:
				return
			default:
				if input.Scan() {
					text <- input.Text()
				}
			}
		}
	}()
selectLoop:
	for {
		select {
		case cur := <-text:
			wg.Add(1)
			count = 0
			go func(shout string) {
				defer wg.Done()
				echo(c, shout, time.Second)
			}(cur)
		case <-ticker.C:
			if count >= 10 {
				done <- struct{}{}
				break selectLoop
			}
			count++
		}
	}

	// NOTE: ignoring potential errors from input.Err()
	wg.Wait()
	fmt.Fprintf(os.Stdout, "connection closed\n")
	c.(*net.TCPConn).CloseWrite()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}

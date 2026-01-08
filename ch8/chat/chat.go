// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// !+broadcaster
type client struct {
	name  string
	cache chan<- string // an outgoing message channel
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.cache <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			messages <- cli.name + " has arrived"

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.cache)
		}
	}
}

//!-broadcaster

// !+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	cli := client{who, ch}
	entering <- cli

	input := bufio.NewScanner(conn)

	reset := make(chan struct{})

	go func() {
		timer := time.NewTimer(5 * time.Minute)
		defer timer.Stop()
	loop:
		for {
			select {
			case <-timer.C:
				conn.Close()
				break loop
			case <-reset:
				if !timer.Stop() {
					select {
					case <-timer.C:
					default:
					}
				}
				timer.Reset(5 * time.Minute)
			}
		}
	}()

	for input.Scan() {
		reset <- struct{}{}
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

// !+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main

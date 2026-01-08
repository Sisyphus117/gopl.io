// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan struct {
	amount   int
	sentback chan<- bool
})

func Deposit(amount int) { deposits <- amount }
func Withdraw(amount int) bool {
	var sentback = make(chan bool)
	withdraws <- struct {
		amount   int
		sentback chan<- bool
	}{amount, sentback}
	return <-sentback
}
func Balance() int { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case msg := <-withdraws:
			if balance >= msg.amount {
				balance -= msg.amount
				msg.sentback <- true
			} else {
				msg.sentback <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-

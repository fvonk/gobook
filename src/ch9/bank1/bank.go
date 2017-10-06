// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

type WithdrawData struct {
	amount int
	result chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan WithdrawData)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdraws <- WithdrawData{amount, ch}
	return <-ch
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdrawData := <- withdraws:
			if balance >= withdrawData.amount {
				balance -= withdrawData.amount
				withdrawData.result <- true
			} else {
				withdrawData.result <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-

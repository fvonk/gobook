// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"testing"

	"ch9/bank1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})
	bank.Deposit(50)

	//go func() {
	//	bank.Withdraw(50)
	//	fmt.Println("=", bank.Balance())
	//	done <- struct{}{}
	//}()

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done
	//<-done

	if got, want := bank.Balance(), 350; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	if got, want := bank.Withdraw(550), false; got != want {
		t.Errorf("Withdraw Balance = %t, want %t", got, want)
	}
	if got, want := bank.Withdraw(150), true; got != want {
		t.Errorf("Withdraw Balance = %t, want %t", got, want)
	}

	if got, want := bank.Balance(), 200; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestBank2(t *testing.T) {
	done := make(chan struct{})
	bank.Withdraw(200)
	// Alice
	go func() {
		bank.Deposit(200)
		if ok := bank.Withdraw(20); ok {
			fmt.Println("Withdraw 20")
		} else {
			fmt.Println("Can't withdraw 20")
		}
		if ok := bank.Withdraw(400); ok {
			fmt.Println("Withdraw 400")
		} else {
			fmt.Println("Can't withdraw 400")
		}
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		if ok := bank.Withdraw(20); ok {
			fmt.Println("Withdraw 20")
		} else {
			fmt.Println("Can't withdraw 20")
		}
		if ok := bank.Withdraw(400); ok {
			fmt.Println("Withdraw 400")
		} else {
			fmt.Println("Can't withdraw 400")
		}
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 260; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
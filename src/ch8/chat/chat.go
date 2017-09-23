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

//!+broadcaster
type client struct {// an outgoing message channel
	ch chan string
	name string
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
			fmt.Println(msg)
			for cli := range clients {
				cli.ch <- msg
			}

		case cli := <-entering:
			cli.ch <- fmt.Sprint("People in chat: ")
			for client := range clients {
				cli.ch <- fmt.Sprintf("%s, ", client.name)
			}
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string)// outgoing client messages
	client := client{ch: ch, name: ""}
	go clientWriter(conn, client.ch)

	enteredNick := false
	const disconnectAfter = 30 * time.Second
	timer := time.After(disconnectAfter)
	end := make(chan struct{})
	scan := make(chan struct{})
	go func() {
		for {
			select {
			case <-timer:
				if enteredNick {
					end <- struct{}{}
				} else {
					conn.Close()
				}
				return
			case <-scan:
				timer = time.After(disconnectAfter)
			}
		}
	}()

	client.ch <- "Enter your nickname"
	input := bufio.NewScanner(conn)
	for input.Scan() {
		scan <- struct{}{}
		nick := input.Text()
		if len(nick) > 1 {
			client.name = nick
			enteredNick = true
			break
		} else {
			client.ch <- "Enter correct nickname"
		}
	}
	if !enteredNick {
		return
	}

	messages <- client.name + " has arrived"
	entering <- client

	go func() {
		input := bufio.NewScanner(conn)
		for input.Scan() {
			messages <- client.name + ": " + input.Text()
			scan <- struct{}{}
		}
		end <- struct{}{}
		return
	}()

loop:
	for {
		select {
		case <- end:
			break loop
		}
	}

	leaving <- client
	messages <- client.name + " has left"
	conn.Close()
}


func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
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

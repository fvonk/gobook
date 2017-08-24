// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/remove", db.remove)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, err := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if err != nil {
		fmt.Fprintf(w, "new price is not correct %q\n", err)
	} else {
		if _, ok := db[item]; ok {
			db[item] = dollars(price)
			fmt.Fprintf(w, "new price %v for %s udpated\n", price, item)
		} else {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
		}
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	newItem := req.URL.Query().Get("item")
	newPrice, err := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if err != nil {
		fmt.Fprintf(w, "new price is not correct %q\n", err)
	} else {
		if _, ok := db[newItem]; ok {
			fmt.Fprintf(w, "new item %v already exists\n", newItem)
		} else {
			db[newItem] = dollars(newPrice)
			fmt.Fprintf(w, "new price %v for %s created\n", newPrice, newItem)
		}

	}
}

func (db database) remove(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "item %s is removed\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
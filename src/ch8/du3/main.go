// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 250.

// The du3 command computes the disk usage of the files in a directory.
package main

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")
type Data struct {
	nfiles int64
	nbytes int64
	root string
}
//!+
func main() {
	// ...determine roots...

	//!-
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//!+
	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan Data)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes, root)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	//!-

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
	data := map[string]Data{}
	//var nfiles, nbytes int64
loop:
	for {
		select {
		case dt, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			if _, ok := data[dt.root]; !ok {
				data[dt.root] = Data{0, 0, dt.root}
			}
			dt.nbytes += data[dt.root].nbytes
			dt.nfiles += data[dt.root].nfiles + 1
			data[dt.root] = dt
			//nfiles++
			//nbytes += size
		case <-tick:
			printDiskUsage(data)
		}
	}

	printDiskUsage(data) // final totals
	//!+
	// ...select loop...
}

//!-

func printDiskUsage(data map[string]Data) {
	//fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
	for root, dt := range data {
		fmt.Printf("%s %d files %.1f GB\n", root, dt.nfiles, float64(dt.nbytes)/1e9)
	}
	fmt.Println()
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
//!+walkDir
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- Data, root string) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes, root)
		} else {
			fileSizes <- Data{0, entry.Size(), root}
		}
	}
}

//!-walkDir

//!+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	// gui token den channel
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return entries

}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subFilePath := filepath.Join(dir, entry.Name())
			go walkDir(subFilePath, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
	// moi khi xong walk dir thi done cai wait group
	// voi moi entry trong entries trong dir thi
	// neu entry la directory
	// wait them 1 cai
	// roi lai walkdir cai dir do
	// ko thi add vao channel
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f KB\n", nfiles, float64(nbytes)/1e3)
}

var verbose = flag.Bool("v", false, "show verbose progress message")

func main() {
	// determine roots
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	// traverse each root of the file tree in parallel
	// tao 1 cai chan de chua file size
	filesizes := make(chan int64)
	var n sync.WaitGroup
	// tao 1 wait group
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, filesizes)
	}

	go func() {
		n.Wait()
		close(filesizes)
	}()

	var nFiles, nSizes int64
	ticker := &time.Ticker{}

	if *verbose {
		ticker = time.NewTicker(500 * time.Millisecond)
	}

loop:
	for {
		select {
		case size, ok := <-filesizes:
			if !ok {
				break loop
				ticker.Stop()
			}
			nFiles++
			nSizes += size
		case <-ticker.C:
			printDiskUsage(nFiles, nSizes)
		}
	}

	printDiskUsage(nFiles, nSizes)
}

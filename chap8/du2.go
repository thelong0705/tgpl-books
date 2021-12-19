package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			path := filepath.Join(dir, entry.Name())
			walkDir(path, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f KB\n", nfiles, float64(nbytes)/1e3)
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v \n", err)
		return nil
	}
	return entries

}

var verbose = flag.Bool("v", false, "show verbose progress message")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	filesizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, filesizes)
		}
		close(filesizes)
	}()
	var nfiles, nsize int64
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
			nfiles++
			nsize += size
		case <-ticker.C:
			printDiskUsage(nfiles, nsize)
		}
	}

	printDiskUsage(nfiles, nsize)
}

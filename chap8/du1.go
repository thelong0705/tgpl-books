package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v \n", err)
		return nil
	}
	return entries

}

func main() {
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}
	filesizes := make(chan int64)
	go func() {
		for _, dir := range roots {
			walkDir(dir, filesizes)
		}
		close(filesizes)
	}()
	var nfiles, nsizes int64

	for filesize := range filesizes {
		nfiles++
		nsizes += filesize
		printDiskUsage(nfiles, nsizes)
	}

}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f KB\n", nfiles, float64(nbytes)/1e3)
}

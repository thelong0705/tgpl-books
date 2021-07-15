package main

import (
	"fmt"
)

func main() {
	fmt.Println([]byte("x"))
	fmt.Println([]byte("X"))
	fmt.Printf("%b", []byte("x"))
	fmt.Println()
	fmt.Printf("%b", []byte("X"))
}

package main

import (
	"fmt"
)


type Currency int
const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main()  {
	symbol := [...]string{RMB: "d", USD: "a", EUR: "b", GBP: "c" }
	fmt.Println(RMB, symbol[RMB])
}

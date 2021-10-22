package main

import "fmt"

func main() {
	f(0)
}

func f(x int) {
	fmt.Printf("f(%d) \n", 0/x)
	defer func() { fmt.Println(1) }()
	defer func() { fmt.Println(2) }()
	defer func() { fmt.Println(3) }()

}

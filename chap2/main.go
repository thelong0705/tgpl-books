package main

import "fmt"

func main()  {
	x := uint64(13)
	for i := 0; i < 4; i++ {
		//fmt.Println(x&1)
		x = x >> 1
	}

	var pc [4]byte

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("%b \n", pc[i])
	}

}

// 5 = 1101 >> 1 -> 0110 >> 1 -> 0011
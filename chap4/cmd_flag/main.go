package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	shaOptPtr := flag.String("sha-opt", "256", "sha option. Its option are 384, 512 otherwise default to 256")
	target := flag.String("target", "", "a string")
	flag.Parse()

	switch shaOpt := *shaOptPtr; shaOpt {
	case "384":
		fmt.Printf("%x \n", sha512.Sum384([]byte(*target)))
	case "512":
		fmt.Printf("%x \n", sha512.Sum512([]byte(*target)))
	default:
		fmt.Printf("%x \n", sha256.Sum256([]byte(*target)))
	}


}
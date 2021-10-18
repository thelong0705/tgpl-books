package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("big slow operation")()
	time.Sleep(time.Second * 5)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("%s", msg)
	return func() { log.Printf("%s end. %s elasped", msg, time.Since(start)) }
}

func main() {
	bigSlowOperation()
}
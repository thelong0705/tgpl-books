package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		_, err := io.Copy(os.Stdout, conn)
		log.Println(err)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}


func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
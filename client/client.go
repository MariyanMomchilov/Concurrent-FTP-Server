package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Fatalln(err)
	}

	ch := make(chan int)

	go func() {
		io.Copy(os.Stdout, conn)
		fmt.Println("Sending to ch...")
		ch <- 0
	}()

	func() {
		io.Copy(conn, os.Stdin)
		fmt.Println("Done...")
	}()

	//!
	conn.Close()
	<-ch
	fmt.Println("Disconnect from server...")

}

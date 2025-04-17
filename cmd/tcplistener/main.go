package main

import (
	"fmt"
	"log"
	"net"
	// "os"

	"github.com/ds1242/httpfromtcp.git/internal/request"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening for TCP traffic: %s\n", err.Error())
	}

	defer listener.Close()

	fmt.Printf("Listening for TCP traffic on : %s\n", port)
	fmt.Println("===============================")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}
		fmt.Println("Connection accepted from ", conn.RemoteAddr())
		go func (c net.Conn) {
			linesChan, err := request.RequestFromReader(c)
			if err != nil {
				fmt.Println("Error reading lines : ", err.Error())
			}

			fmt.Println("Request line:")
			fmt.Printf("- Method: %s\n", linesChan.RequestLine.Method)
			fmt.Printf("- Target: %s\n", linesChan.RequestLine.RequestTarget)
			fmt.Printf("- Version: %s\n", linesChan.RequestLine.HttpVersion)
			fmt.Println("Connection to ", conn.RemoteAddr(), "channel closed")
			c.Close()
		}(conn)
	}
}




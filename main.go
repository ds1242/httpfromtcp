package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	// "os"
	"strings"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening for TCP traffic: %s\n", err.Error())
	}

	defer listener.Close()

	fmt.Printf("Listening on port : %s\n", listener.Addr())
	fmt.Println("===============================")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection accepted")
		go func (c net.Conn) {
			linesChan := getLinesChannel(c)

			for line := range linesChan {
				fmt.Println(line)
			}

			fmt.Println("channel closed")
			c.Close()
		}(conn)
	}
}


func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {
		defer f.Close()
		defer close(ch)
	
		currentLine := ""
		for {

			buffer := make([]byte, 8)
			n, err := f.Read(buffer)
			if err != nil {
				if currentLine != "" {
					ch <- currentLine
					currentLine = ""
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				break
			}

			str := string(buffer[:n])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				ch <- currentLine + parts[i]
				currentLine = ""
			}
			currentLine += parts[len(parts)-1]
		}
	
	}()
	return ch
}

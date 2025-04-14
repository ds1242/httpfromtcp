package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	// "strings"
)

const localhost = "localhost:42069"

func main() {

	address, err := net.ResolveUDPAddr("udp", localhost)
	if err != nil {
		log.Fatalf("error resolving UDP address %s\n", err.Error())
	}

	conn, err := net.DialUDP("udp", nil, address) 
	if err != nil {
		log.Fatalf("error Dialing UDP %s\n", err.Error())
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("> ")
		text, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}
		conn.Write(text)
	}
}
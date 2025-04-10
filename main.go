package main

import (
	"fmt"
	"log"
	"os"
)

const filePath = "messages.txt"

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("could not open %s: %s\n", filePath, err)
	}

	fmt.Printf("Reading data from %s\n", filePath)
	fmt.Println("===============================")

	ch := getLinesChannel(f)
	for line := range ch {
		fmt.Printf("read: %s\n", line)
	}

}



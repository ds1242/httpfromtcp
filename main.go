package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const filePath = "messages.txt"

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("could not open %s: %s\n", filePath, err)
	}

	defer f.Close()

	fmt.Printf("Reading data from %s\n", filePath)
	fmt.Println("===============================")

	var currentLine string
	for {

		b := make([]byte, 8, 8)
		val, err := f.Read(b)
		if err == io.EOF {
			if currentLine != "" {
				fmt.Printf("read: %s\n", currentLine)
			}
			break
		}

		parts := strings.Split(string(b[:val]), "\n")
		if len(parts) > 1 {
			fmt.Printf("read: %s%s\n", currentLine, parts[0])

			for i := 1; i < len(parts)-1; i++ {
				fmt.Printf("read: %s\n", parts[i])
			}

			currentLine = parts[len(parts)-1]
		} else {
			currentLine += parts[0]
		}
	}

}

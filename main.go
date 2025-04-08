package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	b := make([]byte, 8)
	var currentLine string
	for {
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

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
			break
		}

		currentPart := b[:val]
		parts := strings.Split(string(currentPart), "\n")
		// fmt.Println(parts)

		for i := 0; i < len(parts)-1; i++ {
			currentLine += parts[i]
		}
		fmt.Printf("read: %s\n", currentLine)
		currentLine = "" + parts[len(parts)-1]
	}

}

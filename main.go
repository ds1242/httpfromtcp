package main

import (
	"fmt"
	"log"
	"os"
	"errors"
	"strings"
	"io"
)

const filePath = "messages.txt"

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("could not open %s: %s\n", filePath, err)
	}

	fmt.Printf("Reading data from %s\n", filePath)
	fmt.Println("===============================")

	linesChan := getLinesChannel(f)
	for line := range linesChan {
		fmt.Printf("read: %s\n", line)
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

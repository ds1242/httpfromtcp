package main

import (
	"errors"
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

	ch := getLinesChannel(file)
	for line := range ch {
		fmt.Printf("read: %s\n", line)
	}

}

func getLinesChannel(f io.ReadCloser) <-chan string {
	currentLine := ""
	for {

		buffer := make([]byte, 8, 8)
		n, err := f.Read(buffer)
		if err != nil {
			if currentLine != "" {
				fmt.Printf("read: %s\n", currentLine)
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
			fullLine := currentLine + parts[i]
			currentLine = ""
			return fullLine
		}
		currentLine += parts[len(parts)-1]
	}
}

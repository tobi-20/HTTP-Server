package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	const file = `C:\Users\Lawrence Oluwadare's\Desktop\PRIME\messages.txt`
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buff := make([]byte, 8)
	currentLineContents := ""

	for {
		n, err := f.Read(buff)

		if err != nil {
			if currentLineContents != "" {
				fmt.Printf("read: %s\n", currentLineContents)
				currentLineContents = ""
			}
			if err == io.EOF {
				break
			}
			fmt.Printf("error: %s\n", err.Error())
			break
		}

		str := string(buff[:n])
		parts := strings.Split(str, "\n")
		for i := 0; i < len(parts)-1; i++ {
			fmt.Printf("read: %s%s\n", currentLineContents, parts[i])
			currentLineContents = ""
		}
		currentLineContents += parts[len(parts)-1]
	}

}

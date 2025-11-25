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

	linesChan := getLinesChannel(f)

	for line := range linesChan {
		fmt.Println("read:", line)
	}

}
func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {
		defer f.Close()
		defer close(ch)
		currentLineContents := ""
		buff := make([]byte, 8)

		for {
			n, err := f.Read(buff)

			if err != nil {
				if currentLineContents != "" {
					ch <- currentLineContents
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
				ch <- currentLineContents + parts[i]
				currentLineContents = ""
			}
			currentLineContents += parts[len(parts)-1]
		}

	}()
	return ch
}

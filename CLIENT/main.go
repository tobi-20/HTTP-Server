package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:2000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	path := filepath.Join("data", "Renegade.txt")
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(f)

	sentChan := make(chan []byte, 1024)

	go func() {
		defer close(sentChan)
		for {
			lines, err := reader.ReadBytes('\n')

			if err == io.EOF {
				if len(lines) > 0 {
					sentChan <- lines

				}
				return

			}
			if err != nil {
				fmt.Println(err)
			}
			sentChan <- lines

		}
	}()
	for m := range sentChan {

		if _, err := conn.Write(m); err != nil {
			fmt.Println(err)
		}
	}

}

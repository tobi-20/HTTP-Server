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
	l, err := net.Listen("tcp", ":2000") //Listening socket
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	fmt.Println("Server is listening on port 2000")

	//handle multiple clients
	for {
		// This for loop makes possible one server to multiple clients connection (one to many)
		conn, err := l.Accept() //connection socket
		if err != nil {
			fmt.Println(err)
			return
		}

		writeChan := make(chan []byte, 1024) // A buffered channel which sends data when the buffer size is filled up
		reader := bufio.NewReader(conn)

		go handleWrite(conn, writeChan)
		//Read

		go func() {
			defer close(writeChan)
			for {
				lines, err := reader.ReadBytes('\n')
				fmt.Println("server read:", len(lines), "err:", err)
				if err == io.EOF {
					if len(lines) > 0 {
						writeChan <- lines
					}
					return
				}

				if err != nil {
					fmt.Println(err)
					return
				}
				writeChan <- lines // 1. Channels coordinate concurrency between the goroutines
			}
		}()
	}
}
func handleWrite(conn net.Conn, writeChan chan []byte) {
	defer conn.Close()
	outPath := filepath.Join("data", "output.txt")
	f, err := os.Create(outPath)

	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(f)

	defer f.Close()
	// The channel blocks until data is being sent into the channel in  the reading loop see 1. above

	for m := range writeChan {
		if _, err := writer.Write(m); err != nil {
			log.Fatal(err)
		}
		if err := writer.Flush(); err != nil {
			fmt.Println("flush error:", err)
		}
	}

}

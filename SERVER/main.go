package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":2000") //Listening socket
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	fmt.Println("Server is listening on port 2000")

	for {
		conn, err := l.Accept() //connection socket
		if err != nil {
			fmt.Println(err)
			continue
		}

		writeChan := make(chan []byte, 1024)
		go handleWrite(conn, writeChan)
		//Read

		reader := bufio.NewReader(conn)

		go func() {
			defer close(writeChan)
			for {
				lines, err := reader.ReadBytes('\n')
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
				writeChan <- lines
			}
		}()
	}
}
func handleWrite(conn net.Conn, writeChan chan []byte) {
	defer conn.Close()
	for m := range writeChan {
		fmt.Println(string(m))

	}

}

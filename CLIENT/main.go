package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:2000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	sentChan := make(chan []byte, 1024)
	str := "Client here, Hola papi"
	clientMessage := []byte(str)

	sentChan <- clientMessage

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					fmt.Println("Connection ended")
					close(sentChan)
				} else {
					fmt.Println(err)
				}

			}
			received := string(buf[:n])
			fmt.Println("Server replied:", received)

		}

	}()
	for m := range sentChan {

		if _, err := conn.Write(m); err != nil {
			fmt.Println(err)
		}
	}
	conn.Close()
}

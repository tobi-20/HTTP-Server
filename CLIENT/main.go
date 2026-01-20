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

	go func() {
		str := "Client here, Hola papi"
		clientMessage := []byte(str)
		if _, err := conn.Write(clientMessage); err != nil {
			fmt.Println(err)
		}

	}()
	buf := make([]byte, 3)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection ended")
				break
			} else {
				fmt.Println(err)
			}

		}
		received := string(buf[:n])
		fmt.Println(received)

	}
}

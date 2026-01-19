package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	fmt.Println("Server is listening on port 2000")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	str := "Hi this is Lord Gboyega's Server"
	message := []byte(str)
	conn.Write(message)

	defer conn.Close()
}

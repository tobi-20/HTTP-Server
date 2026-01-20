package main

import (
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
		go handleWrite(conn)
		//Read
		buf := make([]byte, 5)
		for {
			if n, err := conn.Read(buf); err != nil {
				if err == io.EOF {
					fmt.Println("End of message")
					break
				}
				fmt.Println(err)
				break
			} else {
				fmt.Println("Client message:", string(buf[:n]))
			}
		}
	}
}
func handleWrite(conn net.Conn) {
	defer conn.Close()

	//Write
	str := "Hi this is Lord Gboyega's Server"
	message := []byte(str)
	conn.Write(message)

}

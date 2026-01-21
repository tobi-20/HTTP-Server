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
		writeChan := make(chan []byte, 1024)
		go handleWrite(conn, writeChan)
		//Read
		buf := make([]byte, 100)

		go func() {
			for {
				if n, err := conn.Read(buf); err != nil {
					if err == io.EOF {
						fmt.Println("End of message")
						close(writeChan)
						break
					}
					fmt.Println(err)
					break
				} else {
					fmt.Println("Client sent:", string(buf[:n]))
				}
			}
		}()
	}
}
func handleWrite(conn net.Conn, writeChan chan []byte) {

	//Write
	str := "Hi this is Lord Gboyega's Server"
	message := []byte(str)
	writeChan <- message

	for m := range writeChan {
		if _, err := conn.Write(m); err != nil {

			fmt.Println(err)
		}

	}
	conn.Close()

}

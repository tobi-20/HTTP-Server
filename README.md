# HTTP-Server
This repository documents my project-based learning journey toward becoming a backend engineer.

The focus of this project is implementing the HTTP protocol from first principles using Go. Since HTTP is an application-layer protocol that runs on TCP, the first step is building and understanding a TCP server from scratch.

The TCP connection model here is many-to-one:
one server handling connections from multiple clients.
The server listens and accepts connections, while clients dial in.

The communication is full-duplex, meaning both the client and server can read and write simultaneously similar to how an email client and email server exchange data independently.

In this implementation, you will see these concepts in practice. Go makes full-duplex communication straightforward through goroutines and channels, which help coordinate concurrent reads and writes while avoiding race conditions.

This project is about understanding how things actually work, not hiding behind frameworks.
Resources:
1. [RFC 793](https://datatracker.ietf.org/doc/html/rfc793#page-1)
2. [Channels in Go by Abu Bakar](https://abubakardev0.medium.com/understanding-channels-in-go-a-comprehensive-guide-a5a9f823c709)
3. [In-depth introduction to bufio.Scanner in Golang](https://medium.com/golangspec/in-depth-introduction-to-bufio-scanner-in-golang-55483bb689b4)
4. [bufio package](https://pkg.go.dev/bufio#pkg-constants)

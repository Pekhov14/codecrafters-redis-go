package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const MAX_BUF_SIZE = 1024

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer listener.Close()

	for {
		// Block until we receive an incoming connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err)
			continue
		}

		// Handle client connection
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// Ensure we close the connection after we're done
	defer conn.Close()

	// Read data
	buf := make([]byte, MAX_BUF_SIZE)
	n, err := conn.Read(buf) // n
	if err != nil {
		return
	}

	log.Println("Received data", buf[:n])

	message := "+PONG\r\n"
	conn.Write([]byte(message))

	// Write the same data back
	// conn.Write(buf[:n])
}

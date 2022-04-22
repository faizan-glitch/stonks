package main

import (
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatal("Server failed to start listening:", err.Error())
	}

	defer lis.Close()

	for {
		conn, err := lis.Accept()
		log.Println("Server accepted connection")

		if err != nil {
			// log.Fatal("Server failed to accept connection:", err.Error())
			log.Println("Server failed to accept connection:", err.Error())
			continue
		}

		go handler(conn)
	}
}

func handler(c net.Conn) {
	buf := make([]byte, 1024)

	defer c.Close()

	_, err := c.Read(buf)

	if err != nil {
		log.Println("Failed to read from connection:", err.Error())
	}

	c.Write([]byte("Hello World!\n"))
}

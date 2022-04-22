package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")

	if err != nil {
		log.Fatal("Failed to connect to server:", err.Error())
	}

	defer conn.Close()

	_, err = conn.Write([]byte("Hello from client!\n"))

	if err != nil {
		log.Fatal("Failed to write to server:", err.Error())
	}

	buf := make([]byte, 1024)

	_, err = conn.Read(buf)

	if err != nil {
		log.Fatal("Failed to read from server:", err.Error())
	}

	log.Println(string(buf))

}

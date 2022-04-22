package main

import (
	"fmt"
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

		fmt.Println(conn)
	}
}

package main

import (
	"encoding/json"
	"log"
	"net"
	"time"

	"github.com/faizan-glitch/stonks/pkg/stocks"
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

	t := time.NewTicker(1 * time.Second)

	defer t.Stop()

	stk := stocks.Stock{
		Time:   time.Now(),
		Symbol: "AAPL",
		Open:   100.0,
		High:   100.0,
		Low:    100.0,
		Close:  100.0,
		Volume: 100,
	}

	for range t.C { // t.C is a channel so we can range over it
		stk.Update()

		b, err := json.MarshalIndent(stk, "", "  ")

		if err != nil {
			log.Println("Failed to marshal stock:", err.Error())
		}

		_, err = c.Write(b)

		if err != nil {
			log.Println("A connection was broken")
			return
		}
	}
}

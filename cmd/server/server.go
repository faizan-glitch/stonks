package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/faizan-glitch/stonks/pkg/cache"
	"github.com/faizan-glitch/stonks/pkg/stocks"
)

var stk = stocks.Stock{
	Time:   time.Now(),
	Symbol: "AAPL",
	Open:   100.0,
	High:   100.0,
	Low:    100.0,
	Close:  100.0,
	Volume: 100,
}

func main() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatal("Server failed to start listening:", err.Error())
	}

	defer lis.Close()

	ch := new(cache.Cache)

	log.Println("Server is running on port 9000")

	for {
		conn, err := lis.Accept()

		log.Println("Server accepted a connection")

		if err != nil {
			// log.Fatal("Server failed to accept connection:", err.Error())
			log.Println("Server failed to accept connection:", err.Error())
			continue
		}

		go handler(conn, ch)
	}
}

func handler(c net.Conn, ch *cache.Cache) {
	buf := make([]byte, 1024)

	defer c.Close()

	_, err := c.Read(buf)

	if err != nil {
		log.Println("Failed to read from connection:", err.Error())
	}

	t := time.NewTicker(1 * time.Second)

	defer t.Stop()

	for range t.C { // t.C is a channel so we can range over it

		if len(ch.Stocks) < cache.Limit {
			fmt.Println("Adding to cache")
			stk.Symbol = stocks.RandomSymbol()
			ch.Add(stk)
		}

		stk = ch.RandomStock()

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

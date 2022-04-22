package cache

import (
	"math/rand"
	"sync"
	"time"

	"github.com/faizan-glitch/stonks/pkg/stocks"
)

// I just realized that Cache doesn't need to be a k-v datastore.
// It can be a struct containing a simple slice

const Limit = 10 // Max number of stocks to keep in the cache

var mut = &sync.Mutex{} // Mutex needed cause multiple goroutines can be trying to add data concurrently

type Cache struct {
	Stocks []stocks.Stock
}

func (c *Cache) Add(s stocks.Stock) {
	if len(c.Stocks) < Limit {
		mut.Lock()
		c.Stocks = append(c.Stocks, s)
		mut.Unlock()
	}
}

func (c *Cache) RandomStock() stocks.Stock {
	rand.Seed(time.Now().Unix())
	// I don't think we need to lock the mutex here. Not sure though.
	return c.Stocks[rand.Intn(len(c.Stocks))]
}

package stocks

import (
	"math/rand"
	"time"
)

type Stock struct {
	Time   time.Time `json:"time"`
	Symbol string    `json:"symbol"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
	Volume int       `json:"volume"`
}

func (s *Stock) Update() {
	rand.Seed(time.Now().Unix())

	s.Time = time.Now()

	s.Symbol = generateRandomSymbol()

	// Todo: Calculate s.Close according to the requirements

	if s.Close > s.High {
		s.High = s.Close
	} else {
		s.Low = s.Close
	}
	s.Volume = rand.Intn(1000)
}

func generateRandomSymbol() string {
	sym := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	output := make([]byte, 4)

	// generate a random symbol of length 4

	for i := 0; i < 4; i++ {
		output[i] = sym[rand.Intn(26)]
	}

	return string(output)
}

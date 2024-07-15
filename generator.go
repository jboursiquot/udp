package udp

import (
	"context"
	"math/rand/v2"
	"time"
)

// Generator generates payloads based on a configuration.
type Generator struct {
	interval int
}

// NewGenerator returns a new Generator.
func NewGenerator(interval int) Generator {
	return Generator{
		interval: interval,
	}
}

// Generate generates Payloads with random data and sends them over the sendChan.
func (g Generator) Generate(ctx context.Context, sendChan chan<- Payload) {
	for {
		ticker := time.NewTicker(time.Second * time.Duration(g.interval))
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			p := Payload{DeviceID: "abc123", Temperature: rand.IntN(100)}
			sendChan <- p
		}
	}
}

package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"time"

	"github.com/jboursiquot/udp"
)

func main() {
	log.Println("resolving...")
	addr, err := net.ResolveUDPAddr("udp", "localhost:9090")
	if err != nil {
		log.Fatalf("failed to resolve UDP address: %s", err)
	}

	log.Println("dialing...")
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("failed to dial: %s", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	receiveChan := make(chan udp.Payload)
	generator := udp.NewGenerator(3)
	go func() {
		log.Println("generating...")
		generator.Generate(ctx, receiveChan)
	}()

	for {
		select {
		case p := <-receiveChan:
			log.Println("<-payloadChan")
			bs, err := json.Marshal(p)
			if err != nil {
				log.Fatalf("failed to marshal payload: %s", err)
			}

			nBytesSent, err := conn.Write(bs)
			if err != nil {
				log.Fatalf("failed to write: %s", err)
			}

			log.Printf("%d bytes sent; %v", nBytesSent, p)
		case <-ctx.Done():
			log.Println("<-ctx.Done()")
			return
		}
	}
}

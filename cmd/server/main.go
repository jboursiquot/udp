package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net"

	"github.com/davecgh/go-spew/spew"
	"github.com/jboursiquot/udp"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:9090")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	log.Printf("Listening on %s", conn.LocalAddr().String())

	buffer := make([]byte, 1024)

	for {
		bytesRead, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatalf("failed to read: %s", err)
		}

		var p udp.Payload
		if err := json.NewDecoder(bytes.NewReader(buffer)).Decode(&p); err != nil {
			log.Fatalf("failed to decode payload: %s", err)
		}

		spew.Dump(bytesRead, clientAddr.String(), p)
	}
}

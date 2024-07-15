package udp

import "fmt"

type Payload struct {
	DeviceID    string `json:"device_id"`
	Temperature int    `json:"temperature"`
}

func (p Payload) String() string {
	return fmt.Sprintf("{device_id: %s, temperature: %d}", p.DeviceID, p.Temperature)
}

package Serial

import (
	"encoding/json"
	"log"

	"go.bug.st/serial"
)

type Messwert struct {
	Ppm int `json:"ppm"`
	Lux int `json:"lux"`

	Hum  float64 `json:"humid"`
	Temp float64 `json:"temp"`
}

type Message struct {
	Sender string `json:"id"`

	Data Messwert `json:"data"`
}

type Arduino struct {
	Id   string
	Port string
}

func ReadData(board *Arduino) Messwert {
	var message Message
	var instream []byte

	mode := &serial.Mode{
		BaudRate: 9600,
	}

	conn, err := serial.Open(board.Port, mode)
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Read(instream)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(instream, &message)

	if message.Sender != board.Id {
		log.Fatal("Contacted Wrong Device")
	}

	return message.Data
}

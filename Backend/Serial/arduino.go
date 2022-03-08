package Serial

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

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

func ReadData(board *Arduino) Message {
	var message Message
	var instream = make([]byte, 100)
	var len int

	mode := &serial.Mode{
		BaudRate: 9600,
		DataBits: 8,
	}

	fmt.Println(board.Port)
	conn, err := serial.Open(board.Port, mode)
	if err != nil {
		log.Fatal(err)
	}

	len, err = conn.Read(instream)
	fmt.Println(len)
	if err != nil {
		log.Fatal(err)
	}

	conn.Close()

	fmt.Println(string(instream))
	json.Unmarshal(instream, &message)
	//fmt.Println(message.Sender)

	if board.Id == "" {
		board.Id = message.Sender
	}

	if message.Sender != board.Id {
		log.Fatal("Contacted Wrong Device")
	}

	return message
}

func LogData(db *sql.DB) {

	ard := GetArduinos()

	for _, a := range ard {
		msg := ReadData(&a)

		fmt.Println(a.Id)
		fmt.Println(a.Port)

		table := msg.Sender
		data := msg.Data

		data_str := fmt.Sprintf("%v, %v, %v, %v ", data.Ppm, data.Hum, data.Temp, data.Lux)

		db_str := "INSERT INTO " + table + "(ppm, humidity, temperature, unused, date, time)\nVALUES (" + data_str + "," + "'" + time.Now().Format("2006-01-02") + " " + time.Now().Format("15:04:05") + "'" + ");"
		/*insert, err := db.Query(db_str)
		if err != nil {
			log.Default()
		}
		insert.Close()*/
		fmt.Println(db_str)
	}

}

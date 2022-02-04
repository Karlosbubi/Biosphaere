package main

import (
	"Biosphaere/Serial"
	"Biosphaere/Server"
)

func main() {
	a := Serial.GetArduinos()
	ReadData(a[len(a)-1])
	Server.RunServer()
}

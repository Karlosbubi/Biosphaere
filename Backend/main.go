package main

import (
	"Biosphaere/Serial"
	"Biosphaere/Server"
)

func main() {
	Serial.PrintPorts()
	Server.RunServer()
}

package main

import (
	"Biosphaere/Serial"
	"Biosphaere/Server"
)

func main() {
	Serial.PrintDetailedPorts()
	Server.RunServer()
}

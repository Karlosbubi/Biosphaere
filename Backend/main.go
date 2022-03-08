package main

import (
	"Biosphaere/Serial"
	"Biosphaere/Server"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:Biosphäre@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	go Serial.LogData(db)
	go Server.RunServer(db)

	for {
		//Let the GO-Routines do their thing... Indefinitely
	}
}

package main

import (
	"assignment-golang-backend/db"
	"log"

	"assignment-golang-backend/server"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Println("Failed to connect DB", err)
		return
	}

	server.Init()
}

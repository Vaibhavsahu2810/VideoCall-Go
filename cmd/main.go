package main

import (
	"log"

	"VideoCall-Go/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}

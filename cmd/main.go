package main

import (
	"log"
	"logging-service-wgo/cmd/internals"
)

func main() {
	log.Fatal(internals.Run())
}

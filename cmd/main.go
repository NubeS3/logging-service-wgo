package main

import (
	"github.com/NubeS3/logging-service-wgo/cmd/internals"
	"log"
)

func main() {
	log.Fatal(internals.Run())
}

package main

import (
	"log"

	"github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/weather-service/internal/infra/di"
)

func main() {
	server := di.NewDi()
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}

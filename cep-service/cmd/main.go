package main

import "github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/cep-service/internal/infra/di"

func main() {
	di := di.NewDI()
	err := di.Start()

	if err != nil {
		panic(err)
	}
}

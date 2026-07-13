package main

import (
	"context"
	"log"
	"time"

	otelpkg "github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/pkg/otel"

	"github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/weather-service/internal/infra/di"
)

func main() {
	ctx := context.Background()

	shutdown, err := otelpkg.SetupOTelSDK("weather-service", ctx)
	if err != nil {
		log.Fatalf("failed to start otel: %v", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := shutdown(ctx); err != nil {
			log.Printf("failed to shutdown OTel: %v", err)
		}
	}()

	server := di.NewDi()
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"context"
	"log"
	"time"

	"github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/cep-service/internal/infra/di"
	otelpkg "github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/pkg/otel"
)

func main() {
	ctx := context.Background()

	shutdown, err := otelpkg.SetupOTelSDK("cep-service", ctx)
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

	container := di.NewDI()
	if err := container.Start(); err != nil {
		log.Fatal(err)
	}
}

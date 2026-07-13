package handler

import "github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/weather-service/internal/ports"

type CLimateHandler struct {
	usecase ports.CLimate
}

func NewClimateHandler(climateUseCase ports.CLimate) *CLimateHandler {
	return &CLimateHandler{
		usecase: climateUseCase,
	}
}

package usecase

import "github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/weather-service/internal/ports"


type ClimateUseCase struct {
	ports ports.CLimate
}

func NewClimateUseCase(ports ports.CLimate) *ClimateUseCase {
	return &ClimateUseCase{
		ports: ports,
	}
}

package ports

import "github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/weather-service/internal/dto"

type CLimate interface {
	GetViaCep(cep string) (*dto.ViaCepResponseOutput, error)
	GetTemp(city string) (*dto.TempResponseOutput, error)
}

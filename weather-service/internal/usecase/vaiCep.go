package usecase

import "github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/weather-service/internal/dto"

func (c *ClimateUseCase) GetViaCep(cep string) (*dto.ViaCepResponseOutput, error) {
	viaCepResponse, err := c.ports.GetViaCep(cep)
	if err != nil {
		return nil, err
	}
	return viaCepResponse, nil
}

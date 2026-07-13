package usecase

import "github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/weather-service/internal/dto"



func (t *ClimateUseCase) GetTemp(city string) (*dto.TempResponseOutput, error) {
	temp, err := t.ports.GetTemp(city)

	if err != nil {
		return nil, err
	}

	return temp, nil
}

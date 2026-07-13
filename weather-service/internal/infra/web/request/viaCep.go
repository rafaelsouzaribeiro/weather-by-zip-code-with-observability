package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/weather-service/internal/dto"
)

var ErrInvalidZipCode = errors.New("invalid zipcode")

var cepRegex = regexp.MustCompile(`^\d{8}$`)

func (r *Request) GetViaCep(cep string) (*dto.ViaCepResponseOutput, error) {
	if !cepRegex.MatchString(cep) {
		return nil, ErrInvalidZipCode
	}

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error in API response: %s", resp.Status)
	}

	var viaCepResponse dto.ViaCepResponseInput
	err = json.NewDecoder(resp.Body).Decode(&viaCepResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}
	return &dto.ViaCepResponseOutput{
		Cep:         viaCepResponse.Cep,
		Logradouro:  viaCepResponse.Logradouro,
		Complemento: viaCepResponse.Complemento,
		Bairro:      viaCepResponse.Bairro,
		Localidade:  viaCepResponse.Localidade,
		Uf:          viaCepResponse.Uf,
		Ibge:        viaCepResponse.Ibge,
		Gia:         viaCepResponse.Gia,
		Ddd:         viaCepResponse.Ddd,
		Siafi:       viaCepResponse.Siafi,
	}, nil
}

package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/cep-service/internal/dto"
)

func (u *UseCase) GetInfo(ctx context.Context, cep string) (*dto.LocaleOuput, error) {
	payload := dto.LocaleInput{Cep: cep}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST",
		fmt.Sprintf("http://%s:%s",
			os.Getenv("WEATHER_SERVICE_SERVER_HOST"),
			os.Getenv("WEATHER_SERVICE_SERVER_PORT")),
		bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output dto.LocaleOuput
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, err
	}
	return &output, nil
}

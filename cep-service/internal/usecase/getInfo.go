package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/cep-service/internal/dto"
)

func (u *UseCase) GetInfo(ctx context.Context, cep string) (*dto.LocaleOuput, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("http://%s:%s/%s",
			os.Getenv("WEATHER_SERVICE_SERVER_HOST"),
			os.Getenv("WEATHER_SERVICE_SERVER_PORT"),
			cep,
		),
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("weather-service retornou status %d: %s", resp.StatusCode, string(b))
	}

	var output dto.LocaleOuput
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, err
	}
	return &output, nil
}

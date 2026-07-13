package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/weather-service/internal/dto"
)

func (r *Request) GetTemp(city string) (*dto.TempResponseOutput, error) {

	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", os.Getenv("KEY_WEATHER_API"), city)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error in API response: %s", resp.Status)
	}

	var tempResponse dto.TempResponseInput
	err = json.NewDecoder(resp.Body).Decode(&tempResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}
	return &dto.TempResponseOutput{
		Currents: tempResponse.Currents,
	}, nil
}

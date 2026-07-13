package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/weather-service/internal/dto"
)

var (
	ErrZipCodeNotFound = errors.New("can not find zipcode")
	ErrTempNotFound    = errors.New("can not find temp")
)

func (h *CLimateHandler) GetClimateByZipCode(w http.ResponseWriter, r *http.Request) {
	cep := r.PathValue("cep")
	viaCep, err := h.usecase.GetViaCep(cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if viaCep.Cep == "" {
		http.Error(w, ErrZipCodeNotFound.Error(), http.StatusNotFound)
		return
	}

	escapedCity := url.QueryEscape(viaCep.Localidade)
	temp, err := h.usecase.GetTemp(escapedCity)
	if err != nil {
		http.Error(w, ErrTempNotFound.Error(), http.StatusNotFound)
		return
	}

	tempF := (temp.Currents.TempC * 1.8) + 32
	tempK := temp.Currents.TempC + 273.15

	response := dto.Current{
		Locale: viaCep.Localidade,
		TempC:  temp.Currents.TempC,
		TempF:  tempF,
		TempK:  tempK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(response)
}

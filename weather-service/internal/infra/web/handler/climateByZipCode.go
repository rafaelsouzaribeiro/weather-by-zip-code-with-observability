package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
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

	response := map[string]float64{
		"temp_C": temp.Currents.TempC,
		"temp_F": tempF,
		"temp_K": tempK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(response)
}

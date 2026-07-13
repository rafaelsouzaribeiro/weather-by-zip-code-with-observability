package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/weather-service/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClimateMock struct {
	mock.Mock
}

func (m *ClimateMock) GetViaCep(cep string) (*dto.ViaCepResponseOutput, error) {
	args := m.Called(cep)
	return args.Get(0).(*dto.ViaCepResponseOutput), args.Error(1)
}

func (m *ClimateMock) GetTemp(city string) (*dto.TempResponseOutput, error) {
	args := m.Called(city)
	return args.Get(0).(*dto.TempResponseOutput), args.Error(1)
}

func TestGetClimateByZipCode_NegativeTemperature_Success(t *testing.T) {
	mockUC := new(ClimateMock)
	h := &CLimateHandler{usecase: mockUC}

	mockUC.On("GetViaCep", "90000000").Return(&dto.ViaCepResponseOutput{
		Cep:        "90000000",
		Localidade: "Porto Alegre",
	}, nil).Once()

	mockUC.On("GetTemp", "Porto+Alegre").Return(&dto.TempResponseOutput{
		Currents: dto.Current{
			TempC: -10.0,
		},
	}, nil).Once()

	req := httptest.NewRequest(http.MethodGet, "/climate/90000000", nil)
	req.SetPathValue("cep", "90000000")
	rr := httptest.NewRecorder()

	h.GetClimateByZipCode(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var body map[string]float64
	err := json.NewDecoder(rr.Body).Decode(&body)
	assert.NoError(t, err)

	assert.Equal(t, -10.0, body["temp_C"])
	assert.InDelta(t, 14.0, body["temp_F"], 0.0001)
	assert.InDelta(t, 263.15, body["temp_K"], 0.0001)

	mockUC.AssertExpectations(t)
}

func TestGetClimateByZipCode_ZeroDegreeTemperature_Success(t *testing.T) {
	mockUC := new(ClimateMock)
	h := &CLimateHandler{usecase: mockUC}

	mockUC.On("GetViaCep", "80000000").Return(&dto.ViaCepResponseOutput{
		Cep:        "80000000",
		Localidade: "Curitiba",
	}, nil).Once()

	mockUC.On("GetTemp", "Curitiba").Return(&dto.TempResponseOutput{
		Currents: dto.Current{
			TempC: 0.0,
		},
	}, nil).Once()

	req := httptest.NewRequest(http.MethodGet, "/climate/80000000", nil)
	req.SetPathValue("cep", "80000000")
	rr := httptest.NewRecorder()

	h.GetClimateByZipCode(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var body map[string]float64
	err := json.NewDecoder(rr.Body).Decode(&body)
	assert.NoError(t, err)

	assert.Equal(t, 0.0, body["temp_C"])
	assert.Equal(t, 32.0, body["temp_F"])
	assert.InDelta(t, 273.15, body["temp_K"], 0.0001)

	mockUC.AssertExpectations(t)
}

func TestGetClimateByZipCode_CityWithSpacesAndSpecialChars_Success(t *testing.T) {
	mockUC := new(ClimateMock)
	h := &CLimateHandler{usecase: mockUC}

	mockUC.On("GetViaCep", "30000000").Return(&dto.ViaCepResponseOutput{
		Cep:        "30000000",
		Localidade: "Belo Horizonte d'Oeste",
	}, nil).Once()

	mockUC.On("GetTemp", "Belo+Horizonte+d%27Oeste").Return(&dto.TempResponseOutput{
		Currents: dto.Current{
			TempC: 20.0,
		},
	}, nil).Once()

	req := httptest.NewRequest(http.MethodGet, "/climate/30000000", nil)
	req.SetPathValue("cep", "30000000")
	rr := httptest.NewRecorder()

	h.GetClimateByZipCode(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var body map[string]float64
	err := json.NewDecoder(rr.Body).Decode(&body)
	assert.NoError(t, err)

	assert.Equal(t, 20.0, body["temp_C"])
	assert.Equal(t, 68.0, body["temp_F"])
	assert.InDelta(t, 293.15, body["temp_K"], 0.0001)

	mockUC.AssertExpectations(t)
}

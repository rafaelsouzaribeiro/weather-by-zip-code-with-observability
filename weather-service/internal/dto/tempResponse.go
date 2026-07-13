package dto

type TempResponseOutput struct {
	Currents Current `json:"current"`
}

type Current struct {
	TempC  float64 `json:"temp_c"`
	TempF  float64 `json:"temp_f"`
	TempK  float64 `json:"temp_k"`
	Locale string  `json:"city"`
}

type TempResponseInput struct {
	Currents Current `json:"current"`
}

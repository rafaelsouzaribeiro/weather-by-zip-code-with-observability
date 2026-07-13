package dto

type LocaleOuput struct {
	Localidade string  `json:"city"`
	TempC      float64 `json:"temp_c"`
	TempF      float64 `json:"temp_f"`
	TempK      float64 `json:"temp_k"`
}

type LocaleInput struct {
	Cep string `json:"cep"`
}

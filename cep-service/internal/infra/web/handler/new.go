package handler

import "github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/cep-service/internal/usecase"

type Handler struct {
	usecase *usecase.UseCase
}

func NewHandler(usecase *usecase.UseCase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

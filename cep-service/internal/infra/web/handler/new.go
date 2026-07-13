package handler

import "github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/cep-service/internal/usecase"

type Handler struct {
	usecase usecase.IUsecase
}

func NewHandler(usecase usecase.IUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

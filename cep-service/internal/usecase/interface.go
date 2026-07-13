package usecase

import (
	"context"

	"github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/cep-service/internal/dto"
)

type IUsecase interface {
	GetInfo(ctx context.Context, cep string) (*dto.LocaleOuput, error)
}

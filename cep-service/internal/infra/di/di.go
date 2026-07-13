package di

import (
	"os"

	"github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/cep-service/internal/infra/web/handler"
	"github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/cep-service/internal/infra/web/server"
	"github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/cep-service/internal/usecase"
)

func NewDI() *server.Server {
	server := server.New(os.Getenv("CEP_SERVICE_SERVER_PORT"))
	usecase := usecase.NewUseCase()
	h := handler.NewHandler(usecase)
	server.AddHandler("POST {cep}", h.ForwardCep)
	return server
}

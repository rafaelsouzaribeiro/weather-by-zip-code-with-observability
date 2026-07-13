package di

import (
	"os"

	"github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/weather-service/internal/infra/web/handler"
	"github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/weather-service/internal/infra/web/request"
	"github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/weather-service/internal/infra/web/server"
	"github.com/rafaelsouzaribeiro/weather-by-zip-code-with-observability/weather-service/internal/usecase"
)

func NewDi() *server.Server {
	server := server.New(os.Getenv("WEATHER_SERVICE_SERVER_PORT"))
	request := request.NewRequest()
	usecases := usecase.NewClimateUseCase(request)
	handler := handler.NewClimateHandler(usecases)
	server.AddHandler("GET /{cep}", handler.GetClimateByZipCode)

	return server
}

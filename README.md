# Weather by Zip Code with Observability

Sistema distribuído em Go com dois microsserviços para consulta de clima por CEP, com **observabilidade** usando **OpenTelemetry (OTEL)** e **Zipkin** para rastreamento distribuído.

## Arquitetura

- **Serviço A (cep-service)**  
  Recebe a requisição do usuário, valida o CEP e encaminha para o Serviço B.

- **Serviço B (weather-service)**  
  Recebe o CEP, identifica a cidade, consulta a temperatura na WeatherAPI e retorna:
  - Temperatura em Celsius
  - Conversão para Fahrenheit
  - Conversão para Kelvin

- **Observabilidade (OTEL + Zipkin)**  
  Coleta e visualiza traces distribuídos entre os dois serviços.

---

## Pré-requisitos

- Docker e Docker Compose instalados
- Go (opcional, caso rode local sem container)
- Chave da WeatherAPI

---

## 1) Criar conta na WeatherAPI

1. Acesse: [https://www.weatherapi.com/](https://www.weatherapi.com/)
2. Crie uma conta gratuita
3. Gere uma chave de API
4. Coloque a chave no `docker-compose.yaml`:

```yaml
KEY_WEATHER_API=SUA_CHAVE_AQUI
```

---

## 2) Configuração de ambiente

No arquivo `docker-compose.yaml`, ajuste:

```yaml
- KEY_WEATHER_API=SUA_CHAVE_AQUI
```


## 3) Subir o ambiente

Na raiz do projeto:

```bash
docker compose up --build
```

Serviços expostos:

- **cep-service**: `http://localhost:8080`
- **weather-service**: `http://localhost:8081`
- **zipkin**: `http://localhost:9411`

---

## 4) Como testar

Exemplo de requisição para o Serviço A:

```http
POST http://localhost:8080/29902555 HTTP/1.1
Host: localhost:8080
```

> O `cep-service` recebe o CEP e orquestra a chamada ao `weather-service`.

---

## 5) Observabilidade com Zipkin

1. Acesse: [http://localhost:9411](http://localhost:9411)
2. Clique em **Run Query**
3. Visualize os spans dos serviços:
   - `cep-service`
   - `weather-service`

Isso permite acompanhar a requisição ponta a ponta (distributed tracing).

---

## Estrutura (resumo)

- `cep-service/` → Serviço A
- `weather-service/` → Serviço B
- `pkg/otel/` → Inicialização OTEL (trace, metric, log)
- `docker-compose.yaml` → Orquestração local com Zipkin

---

## Observações

- Se `KEY_WEATHER_API` estiver vazia, o `weather-service` pode falhar ao consultar clima.
- Se houver erro de método (ex.: `Method Not Allowed`), valide se a rota esperada é `POST /{cep}`.
- Para desligar:
  ```bash
  docker compose down
  ```
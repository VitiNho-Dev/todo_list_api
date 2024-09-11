# Etapa de compilação
FROM golang:1.23-alpine AS build

WORKDIR /app

# Copia o go.mod e o go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copia o código-fonte
COPY . .

# Compila o binário
RUN go build -o main ./cmd/api/main.go

# Etapa de produção
FROM alpine:3.18

WORKDIR /root/

# Copia o binário da etapa de compilação
COPY --from=build /app/main .

# Define a porta padrão
EXPOSE 8080

# Executa a aplicação
CMD ["./main"]
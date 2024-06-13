FROM golang:1.22-alpine as builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia o código fonte para o diretório de trabalho
COPY . .

# Compila o código Go
RUN go mod tidy
RUN go build -o payroll-api ./cmd/main.go

# Cria uma imagem mínima para execução da aplicação
FROM alpine:latest

# Define o diretório de trabalho dentro do container
WORKDIR /root/

# Copia o executável compilado da fase anterior
COPY --from=builder /app/payroll-api .

# Expõe a porta em que a aplicação Go irá rodar
EXPOSE 8080

# Comando para iniciar a aplicação quando o container for iniciado
CMD ["./payroll-api"]

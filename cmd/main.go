// /cmd/main.go
package main

import (
	"https://github.com/dennersousa/payroll-api/internal/api"
)

func main() {
	router := api.SetupRouter()
	router.Run(":8080") // Inicia o servidor na porta 8080
}

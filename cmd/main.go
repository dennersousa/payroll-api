// cmd/main.go

package main

import (
	"github.com/gin-gonic/gin"
	"payroll-api/internal/api/handlers"
)

func main() {
	r := gin.Default()
	r.POST("/api/calcular", handlers.CalcularFolhaPagamento)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

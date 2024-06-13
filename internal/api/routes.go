// /internal/api/routes.go
package api

import (
	"github.com/gin-gonic/gin"
	"payroll-api/internal/api/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Adicionar middleware se necessário

	api := router.Group("/api")
	{
		api.POST("/calcular", handlers.CalcularFolhaPagamento)
	}

	return router
}

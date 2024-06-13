// internal/api/handlers/calculations.go

package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payroll-api/pkg/calculations"
)

type CalculoRequest struct {
	Colaborador    calculations.Colaborador `json:"colaborador"`
	InicioPeriodo  string                   `json:"inicio_periodo"`
	FimPeriodo     string                   `json:"fim_periodo"`
	JornadaSemanal int                      `json:"jornada_semanal"`
}

func CalcularFolhaPagamento(c *gin.Context) {
	var req CalculoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado, err := calculations.CalcularSalario(req.Colaborador, req.InicioPeriodo, req.FimPeriodo, req.JornadaSemanal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultado)
}

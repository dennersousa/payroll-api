// /internal/api/handlers/calculations_handler.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/payroll-api/pkg/calculations"
)

func CalculatePayroll(c *gin.Context) {
	var input calculations.PayrollInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := calculations.CalculatePayroll(input)
	c.JSON(http.StatusOK, result)
}

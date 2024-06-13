// /pkg/calculations/irrf.go

package calculations

const (
	faixa1MaxIRRF = 2112.00
	faixa2MaxIRRF = 2826.66
	faixa3MaxIRRF = 3751.06
	faixa4MaxIRRF = 4664.68
)

func CalcularIRRF(salarioBruto, inss float64) float64 {
	baseCalculo := salarioBruto - inss

	var irrf float64
	switch {
	case baseCalculo <= faixa1MaxIRRF:
		irrf = 0
	case baseCalculo <= faixa2MaxIRRF:
		irrf = (baseCalculo - faixa1MaxIRRF) * 0.075
	case baseCalculo <= faixa3MaxIRRF:
		irrf = (baseCalculo-faixa2MaxIRRF)*0.15 + (faixa2MaxIRRF-faixa1MaxIRRF)*0.075
	case baseCalculo <= faixa4MaxIRRF:
		irrf = (baseCalculo-faixa3MaxIRRF)*0.225 + (faixa3MaxIRRF-faixa2MaxIRRF)*0.15 + (faixa2MaxIRRF-faixa1MaxIRRF)*0.075
	default:
		irrf = (baseCalculo-faixa4MaxIRRF)*0.275 + (faixa4MaxIRRF-faixa3MaxIRRF)*0.225 + (faixa3MaxIRRF-faixa2MaxIRRF)*0.15 + (faixa2MaxIRRF-faixa1MaxIRRF)*0.075
	}

	if irrf < 0 {
		irrf = 0
	}
	return irrf
}

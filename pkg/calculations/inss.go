// /pkg/calculations/inss.go
package calculations

const (
	salarioMinimo = 1412
	faixa1MaxINSS = 1412
	faixa2MaxINSS = 2666.68
	faixa3MaxINSS = 4000.03
	faixa4MaxINSS = 7786.02
)

func CalcularINSS(salarioBruto float64) float64 {
	switch {
	case salarioBruto <= faixa1MaxINSS:
		return salarioBruto * 0.075
	case salarioBruto <= faixa2MaxINSS:
		return faixa1MaxINSS*0.075 + (salarioBruto-faixa1MaxINSS)*0.09
	case salarioBruto <= faixa3MaxINSS:
		return faixa1MaxINSS*0.075 + (faixa2MaxINSS-faixa1MaxINSS)*0.09 + (salarioBruto-faixa2MaxINSS)*0.12
	case salarioBruto <= faixa4MaxINSS:
		return faixa1MaxINSS*0.075 + (faixa2MaxINSS-faixa1MaxINSS)*0.09 + (faixa3MaxINSS-faixa2MaxINSS)*0.12 + (salarioBruto-faixa3MaxINSS)*0.14
	default:
		return faixa1MaxINSS*0.075 + (faixa2MaxINSS-faixa1MaxINSS)*0.09 + (faixa3MaxINSS-faixa2MaxINSS)*0.12 + (faixa4MaxINSS-faixa3MaxINSS)*0.14
	}
}

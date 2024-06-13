// pkg/calculations/salario.go

package calculations

type Colaborador struct {
	Nome                    string  `json:"nome"`
	SalarioMensal           float64 `json:"salario_mensal"`
	HorasTrabalhadas        float64 `json:"horas_trabalhadas"`
	HorasExtrasDiurnas      float64 `json:"horas_extras_diurnas"`
	HorasExtrasNoturnas     float64 `json:"horas_extras_noturnas"`
	HorasExtrasDomFer       float64 `json:"horas_extras_dom_feriado"`
	HorasNoturnas           float64 `json:"horas_noturnas"`
	AdicionalPericulosidade bool    `json:"adicional_periculosidade"`
}

type ResultadoCalculo struct {
	SalarioBase                float64 `json:"salario_base"`
	Adicional                  float64 `json:"adicional"`
	SalarioTotal               float64 `json:"salario_total"`
	SalarioHora                float64 `json:"salario_hora"`
	AdicionalNoturno           float64 `json:"adicional_noturno"`
	HoraExtraDiurna            float64 `json:"hora_extra_diurna"`
	HoraExtraNoturna           float64 `json:"hora_extra_noturna"`
	HoraExtraDomFer            float64 `json:"hora_extra_dom_fer"`
	ReflexoHorasExtrasDSR      float64 `json:"reflexo_horas_extras_dsr"`
	ReflexoAdicionalNoturnoDSR float64 `json:"reflexo_adicional_noturno_dsr"`
	TotalDevido                float64 `json:"total_devido"`
	INSS                       float64 `json:"inss"`
	IRRF                       float64 `json:"irrf"`
	TotalPagar                 float64 `json:"total_pagar"`
}

func CalcularSalario(colaborador Colaborador, inicioPeriodo, fimPeriodo string, jornadaSemanal int) (ResultadoCalculo, error) {
	_, err := NovoPeriodoTrabalhado(inicioPeriodo, fimPeriodo)
	if err != nil {
		return ResultadoCalculo{}, err
	}

	divisor, err := ObterDivisor(jornadaSemanal)
	if err != nil {
		return ResultadoCalculo{}, err
	}

	salarioBase := colaborador.SalarioMensal
	var adicional float64

	if colaborador.AdicionalPericulosidade {
		adicional = salarioBase * 0.3
	}

	salarioTotal := salarioBase + adicional
	salarioHora := salarioTotal / float64(divisor)

	adicionalNoturno, horaExtraDiurna, horaExtraNoturna, horaExtraDomFer, reflexoHorasExtrasDSR, reflexoAdicionalNoturnoDSR := CalcularHorasExtrasEAdicionais(
		salarioHora, colaborador.HorasExtrasDiurnas, colaborador.HorasExtrasNoturnas, colaborador.HorasExtrasDomFer, colaborador.HorasNoturnas)

	totalDevido := salarioTotal + adicionalNoturno + horaExtraDiurna + horaExtraNoturna + horaExtraDomFer + reflexoHorasExtrasDSR + reflexoAdicionalNoturnoDSR

	inss := CalcularINSS(totalDevido)
	irrf := CalcularIRRF(totalDevido, inss)

	totalPagar := totalDevido - inss - irrf

	return ResultadoCalculo{
		SalarioBase:                salarioBase,
		Adicional:                  adicional,
		SalarioTotal:               salarioTotal,
		SalarioHora:                salarioHora,
		AdicionalNoturno:           adicionalNoturno,
		HoraExtraDiurna:            horaExtraDiurna,
		HoraExtraNoturna:           horaExtraNoturna,
		HoraExtraDomFer:            horaExtraDomFer,
		ReflexoHorasExtrasDSR:      reflexoHorasExtrasDSR,
		ReflexoAdicionalNoturnoDSR: reflexoAdicionalNoturnoDSR,
		TotalDevido:                totalDevido,
		INSS:                       inss,
		IRRF:                       irrf,
		TotalPagar:                 totalPagar,
	}, nil
}

func CalcularHorasExtrasEAdicionais(
	salarioHora float64, horasExtrasDiurnas, horasExtrasNoturnas, horasExtrasDomFer, horasNoturnas float64,
) (float64, float64, float64, float64, float64, float64) {
	adicionalNoturno := salarioHora * 0.20 * horasNoturnas
	horaExtraDiurna := salarioHora * 1.50 * horasExtrasDiurnas
	horaExtraNoturna := salarioHora * 1.80 * horasExtrasNoturnas
	horaExtraDomFer := salarioHora * 2.00 * horasExtrasDomFer

	reflexoHorasExtrasDSR := (horaExtraDiurna + horaExtraNoturna + horaExtraDomFer) / 6.00
	reflexoAdicionalNoturnoDSR := adicionalNoturno / 6.00

	return adicionalNoturno, horaExtraDiurna, horaExtraNoturna, horaExtraDomFer, reflexoHorasExtrasDSR, reflexoAdicionalNoturnoDSR
}

// pkg/calculations/data.go

package calculations

import (
	"errors"
	"time"
)

type PeriodoTrabalhado struct {
	Inicio time.Time
	Fim    time.Time
}

func NovoPeriodoTrabalhado(inicio, fim string) (PeriodoTrabalhado, error) {
	layout := "02/01/2006"
	inicioData, err := time.Parse(layout, inicio)
	if err != nil {
		return PeriodoTrabalhado{}, err
	}
	fimData, err := time.Parse(layout, fim)
	if err != nil {
		return PeriodoTrabalhado{}, err
	}
	if inicioData.After(fimData) {
		return PeriodoTrabalhado{}, errors.New("data de início é posterior à data de fim")
	}
	return PeriodoTrabalhado{Inicio: inicioData, Fim: fimData}, nil
}

func DiasUteisNoMes(periodo PeriodoTrabalhado) int {
	diasUteis := 0
	for dia := periodo.Inicio; !dia.After(periodo.Fim); dia = dia.AddDate(0, 0, 1) {
		if dia.Weekday() != time.Saturday && dia.Weekday() != time.Sunday {
			diasUteis++
		}
	}
	return diasUteis
}

func QuantidadeDeRepousos(periodo PeriodoTrabalhado) int {
	totalRepousos := 0
	for dia := periodo.Inicio; !dia.After(periodo.Fim); dia = dia.AddDate(0, 0, 1) {
		if dia.Weekday() == time.Saturday || dia.Weekday() == time.Sunday {
			totalRepousos++
		}
	}
	return totalRepousos
}

func ObterDivisor(jornadaSemanal int) (int, error) {
	switch jornadaSemanal {
	case 44:
		return 220, nil
	case 40:
		return 200, nil
	case 36:
		return 180, nil
	case 30:
		return 150, nil
	default:
		return 0, errors.New("jornada semanal inválida")
	}
}

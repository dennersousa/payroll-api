// /pkg/calculations/calculations.go
package calculations

type PayrollInput struct {
	Salary    float64 `json:"salary"`
	WorkHours int     `json:"work_hours"`
	Overtime  int     `json:"overtime"`
}

type PayrollOutput struct {
	GrossSalary   float64 `json:"gross_salary"`
	OvertimePay   float64 `json:"overtime_pay"`
	NetSalary     float64 `json:"net_salary"`
	TaxDeductions float64 `json:"tax_deductions"`
}

func CalculatePayroll(input PayrollInput) PayrollOutput {
	hourlyRate := input.Salary / float64(input.WorkHours)
	overtimePay := hourlyRate * 1.5 * float64(input.Overtime)
	grossSalary := input.Salary + overtimePay
	taxDeductions := grossSalary * 0.15 // Exemplo de dedução de 15%
	netSalary := grossSalary - taxDeductions

	return PayrollOutput{
		GrossSalary:   grossSalary,
		OvertimePay:   overtimePay,
		NetSalary:     netSalary,
		TaxDeductions: taxDeductions,
	}
}

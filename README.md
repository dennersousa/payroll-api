# Payroll API

## Descrição

A Payroll API é um projeto desenvolvido em Go, utilizando o framework GinGonic, para realizar cálculos trabalhistas, incluindo remuneração mensal, horas extras, adicional noturno e adicional de periculosidade. Esta API é projetada para calcular a folha de pagamento com base nas regras trabalhistas vigentes na CLT (Consolidação das Leis do Trabalho) do Brasil.

O objetivo principal deste projeto é auxiliar pequenas e médias empresas a lidar com a complexidade dos cálculos da folha de pagamento, reduzindo a necessidade de contratar contadores e minimizando os riscos de erros e processos trabalhistas.

Com essa API, é possível automatizar e garantir a precisão dos cálculos de folha de pagamento, economizando tempo e minimizando erros. A flexibilidade da linguagem Go permite que a API seja escalável e eficiente, proporcionando um desempenho rápido e confiável.

Além disso, a API pode ser integrada diretamente com o Excel, permitindo que os resultados dos cálculos sejam recebidos diretamente nas células do Excel, facilitando a geração das folhas de pagamento.

## Estrutura do Projeto

A estrutura do projeto é organizada conforme abaixo:

```
/payroll-api
|-- /cmd
----|-- main.go
|-- /internal
----|--/api
-------|-- /handlers
-----------|-- calculations.go
-------/middleware
------- routes.go
|-- /pkg
----|-- /calculations
--------|-- data.go
--------|-- inss.go
--------|-- irrf.go
--------|-- salario.go
|-- go.mod
```

### Descrição dos Diretórios e Arquivos

- **/cmd**: Contém o ponto de entrada da aplicação.
    - `main.go`: Inicializa o servidor e define a rota principal da API.

- **/internal**: Contém a lógica interna da aplicação.
    - **/api**
        - **/handlers**
            - `calculations.go`: Implementa o handler para a rota de cálculo da folha de pagamento.
        - `routes.go`: Configura as rotas da API.
    - **/middleware**: (Atualmente vazio, mas preparado para adição de middlewares no futuro).

- **/pkg**: Contém bibliotecas e pacotes reutilizáveis.
    - **/calculations**
        - `data.go`: Define estruturas e funções relacionadas ao período trabalhado.
        - `inss.go`: Contém a lógica para cálculo do INSS.
        - `irrf.go`: Contém a lógica para cálculo do IRRF.
        - `salario.go`: Contém a lógica para cálculo de salário e adicionais.

## Funcionalidades

- Cálculo do salário base, adicionais e total devido.
- Cálculo de horas extras (diurnas, noturnas e em domingos/feriados).
- Cálculo de adicional noturno.
- Cálculo de reflexo das horas extras e adicional noturno no Descanso Semanal Remunerado (DSR).
- Cálculo dos descontos de INSS e IRRF.
- Retorno do valor total a pagar ao empregado.

## Estrutura da API

### Endpoint

#### POST /api/calcular

Calcula a folha de pagamento de um empregado para um determinado período.

**Request:**

```json
{
  "colaborador": {
    "nome": "João da Silva",
    "salario_mensal": 3000.00,
    "horas_trabalhadas": 160,
    "horas_extras_diurnas": 10,
    "horas_extras_noturnas": 5,
    "horas_extras_dom_feriado": 2,
    "horas_noturnas": 15,
    "adicional_periculosidade": true
  },
  "inicio_periodo": "01/06/2024",
  "fim_periodo": "30/06/2024",
  "jornada_semanal": 44
}
```

**Response:**

```json
{
  "salario_base": 3000,
  "adicional": 900,
  "salario_total": 3900,
  "salario_hora": 17.727272727272727,
  "adicional_noturno": 53.18181818181818,
  "hora_extra_diurna": 265.9090909090909,
  "hora_extra_noturna": 159.54545454545456,
  "hora_extra_dom_fer": 70.9090909090909,
  "reflexo_horas_extras_dsr": 82.72727272727273,
  "reflexo_adicional_noturno_dsr": 8.863636363636363,
  "total_devido": 4541.136363636364,
  "inss": 454.5780909090909,
  "irrf": 267.7466113636364,
  "total_pagar": 3818.8116613636366
}
```

## Como Executar

### Passos para executar

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/payroll-api.git
cd payroll-api
```

2. Instale as dependências:

```bash
go mod tidy
```

3. Execute a aplicação:

```bash
go run cmd/main.go
```

A API estará disponível em `http://localhost:8080`.

## Estruturas e Funções Detalhadas

### Estrutura `Colaborador`

Define os dados do empregado necessários para o cálculo da folha de pagamento.

```go
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
```

### Função `CalcularSalario`

Realiza o cálculo completo da folha de pagamento.

```go
func CalcularSalario(colaborador Colaborador, inicioPeriodo, fimPeriodo string, jornadaSemanal int) (ResultadoCalculo, error)
```

### Função `CalcularINSS`

Calcula o desconto de INSS baseado no salário bruto.

```go
func CalcularINSS(salarioBruto float64) float64
```

### Função `CalcularIRRF`

Calcula o desconto de IRRF baseado no salário bruto e no valor do INSS.

```go
func CalcularIRRF(salarioBruto, inss float64) float64
```

### Função `CalcularHorasExtrasEAdicionais`

Calcula o valor das horas extras e adicionais noturnos.

```go
func CalcularHorasExtrasEAdicionais(
    salarioHora float64, horasExtrasDiurnas, horasExtrasNoturnas, horasExtrasDomFer, horasNoturnas float64,
) (float64, float64, float64, float64, float64, float64)
```

## Integração com Excel
A API pode ser integrada diretamente com o Excel, permitindo que os resultados dos cálculos sejam recebidos diretamente nas células do Excel. Isso facilita a geração das folhas de pagamento, economizando tempo e garantindo a precisão dos cálculos.

Para integrar com o Excel, você pode usar o recurso de consulta de dados externos do Excel e configurar uma conexão HTTP para fazer chamadas à API. Assim, você pode preencher automaticamente as células com os resultados dos cálculos trabalhistas.

## Licença

Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.

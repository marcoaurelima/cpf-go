package cpf

import (
	"crypto/rand"
)

var pesos = [10]int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

type CPF struct {
	digits [11]byte
}

// New é a função construtora para criar um novo objeto CPF.
// Ela recebe uma string sem formatação (apenas dígitos numéricos) representando o CPF,
// valida a entrada e retorna um objeto CPF ou um erro caso a validação falhe.
func New(digits string) (CPF, error) {
	if err := validateInput(digits); err != nil {
		return CPF{}, err
	}

	cpf := CPF{
		digits: stringToBytes(digits),
	}

	if err := cpf.isValid(); err != nil {
		return CPF{}, err
	}

	return cpf, nil
}

// Base retorna a base do CPF, ou seja, os primeiros 9 dígitos
func (c CPF) Base() string {
	return string(c.digits[:9])
}

// DV retorna os dígitos verificadores do CPF
func (c CPF) DV() string {
	return string(c.digits[9:11])
}

// String retorna o CPF no formato sem formatação, ou seja, apenas os dígitos
func (c CPF) String() string {
	return string(c.digits[:])
}

// StringFormatted retorna o CPF formatado no padrão XXX.XXX.XXX-XX
func (c CPF) StringFormatted() string {
	return string(c.digits[:3]) + "." + string(c.digits[3:6]) + "." + string(c.digits[6:9]) + "-" + string(c.digits[9:11])
}

// IsValid verifica se o CPF é válido, retornando true se for válido e false caso contrário
func (c CPF) IsValid() bool {
	return (c.isValid() == nil)
}

// isValid realiza a validação do CPF, verificando se os dígitos verificadores estão corretos
func (c CPF) isValid() error {
	// Cálculo dos dígitos verificadores
	dv1, dv2 := calculateDV([9]byte(c.digits[:9]))

	// Comparação dos dígitos verificadores calculados com os fornecidos
	if (dv1 != int(c.digits[9]-'0')) || (dv2 != int(c.digits[10]-'0')) {
		return ErrorCPFInvalido
	}

	return nil
}

// calculateDV é uma função auxiliar que calcula os dígitos verificadores do CPF com base nos primeiros 9 dígitos
func calculateDV(base [9]byte) (dv1, dv2 int) {
	sum1 := 0
	sum2 := 0

	for i := 0; i < len(base); i++ {
		sum1 += int(base[i]-'0') * pesos[i+1]
		sum2 += int(base[i]-'0') * pesos[i]
	}

	// Cálculo do primeiro dígito verificador
	dv1 = (sum1 * 10) % 11
	if dv1 == 10 {
		dv1 = 0
	}

	// Cálculo do segundo dígito verificador
	sum2 += (dv1) * pesos[len(base)]
	dv2 = (sum2 * 10) % 11
	if dv2 == 10 {
		dv2 = 0
	}

	return
}

// NewRandom é uma função que gera um CPF aleatório válido.
func NewRandom() (CPF, error) {
	limit := 1000 // Limite de tentativas para evitar loops infinitos

	var randomDigits [11]byte
	for i := 0; i < limit; i++ {
		for i := range randomDigits {
			d := [1]byte{}
			rand.Read(d[:])
			randomDigits[i] = (d[0] % 10) + '0'
		}

		// Calcular os dígitos verificadores para o CPF gerado
		dv1, dv2 := calculateDV([9]byte(randomDigits[:9]))
		randomDigits[9] = byte(dv1) + '0'
		randomDigits[10] = byte(dv2) + '0'

		cpf, err := New(string(randomDigits[:]))
		if err != nil {
			continue // Se ocorrer um erro, tente gerar outro CPF
		}

		if cpf.IsValid() {
			return cpf, nil
		}
	}

	return CPF{}, ErrorGerarCPFAleatorio
}

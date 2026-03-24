package cpf

import "fmt"

var pesos = [11]int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

type CPF struct {
	base string
	dv   string
}

// New é a função construtora para criar um novo objeto CPF.
// Ela recebe uma string sem formatação (apenas dígitos numéricos) representando o CPF,
// valida a entrada e retorna um objeto CPF ou um erro caso a validação falhe.
func New(cpf string) (CPF, error) {
	if err := validateInput(cpf); err != nil {
		return CPF{}, err
	}

	newCPF := CPF{
		base: cpf[:9],
		dv:   cpf[9:],
	}

	if err := newCPF.isValid(); err != nil {
		return CPF{}, err
	}

	return newCPF, nil
}

// IsValid verifica se o CPF é válido, retornando true se for válido e false caso contrário
func (c CPF) IsValid() bool {
	return (c.isValid() == nil)
}

// Base retorna a base do CPF, ou seja, os primeiros 9 dígitos
func (c CPF) Base() string {
	return c.base
}

// DV retorna os dígitos verificadores do CPF
func (c CPF) DV() string {
	return c.dv
}

// String retorna o CPF no formato sem formatação, ou seja, apenas os dígitos
func (c CPF) String() string {
	return fmt.Sprintf("%s%s", c.base, c.dv)
}

// StringFormatted retorna o CPF formatado no padrão XXX.XXX.XXX-XX
func (c CPF) StringFormatted() string {
	if err := validateInput(c.String()); err != nil {
		return c.String()
	}

	return fmt.Sprintf("%s.%s.%s-%s", c.base[:3], c.base[3:6], c.base[6:9], c.dv)
}

// isValid realiza a validação do CPF, verificando se os dígitos verificadores estão corretos
func (c CPF) isValid() error {
	if err := validateInput(c.String()); err != nil {
		return err
	}

	soma1 := 0
	soma2 := 0

	for i := 0; i < len(c.base); i++ {
		soma1 += int(c.base[i]-'0') * pesos[i+1]
		soma2 += int(c.base[i]-'0') * pesos[i]
	}

	// Cálculo do primeiro dígito verificador
	dv1 := (soma1 * 10) % 11
	if dv1 == 10 {
		dv1 = 0
	}

	// Cálculo do segundo dígito verificador
	soma2 += (dv1) * pesos[len(c.base)]
	dv2 := (soma2 * 10) % 11
	if dv2 == 10 {
		dv2 = 0
	}

	// Comparação dos dígitos verificadores calculados com os fornecidos
	cmp := (dv1 == int(c.dv[0]-'0') && dv2 == int(c.dv[1]-'0'))
	if !cmp {
		return ErrorCPFInvalido
	}

	return nil
}

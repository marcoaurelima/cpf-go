package cpf

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
	soma1 := 0
	soma2 := 0

	for i := 0; i < len(c.digits[:9]); i++ {
		soma1 += int(c.digits[i]-'0') * pesos[i+1]
		soma2 += int(c.digits[i]-'0') * pesos[i]
	}

	// Cálculo do primeiro dígito verificador
	dv1 := (soma1 * 10) % 11
	if dv1 == 10 {
		dv1 = 0
	}

	// Cálculo do segundo dígito verificador
	soma2 += (dv1) * pesos[len(c.digits[:9])]
	dv2 := (soma2 * 10) % 11
	if dv2 == 10 {
		dv2 = 0
	}

	// Comparação dos dígitos verificadores calculados com os fornecidos
	cmp := (dv1 == int(c.digits[9]-'0') && dv2 == int(c.digits[10]-'0'))
	if !cmp {
		return ErrorCPFInvalido
	}

	return nil
}

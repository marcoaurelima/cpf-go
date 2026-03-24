package cpf

// validateInput é uma função auxiliar que valida a entrada do CPF, verificando se a string está vazia,
func validateInput(input string) error {
	if input == "" {
		return ErrorCPFVazio
	}

	if len(input) != 11 {
		return ErrorCPFIncompleto
	}

	if !isNumeric(input) {
		return ErrorCPFNaoNumerico
	}

	if allDigitsEqual(input) {
		return ErrorCPFInvalido
	}

	return nil
}

func allDigitsEqual(input string) bool {
	for i := 1; i < len(input); i++ {
		if input[i] != input[0] {
			return false
		}
	}

	return true
}

// isNumeric é uma função auxiliar que verifica se uma string contém apenas caracteres numéricos
func isNumeric(input string) bool {
	for _, char := range input {
		if char < '0' || char > '9' {
			return false
		}
	}

	return true
}

// stringToBytes é uma função auxiliar que converte uma string de dígitos numéricos em um array de bytes
func stringToBytes(input string) (digits [11]byte) {
	copy(digits[:], input)
	return
}

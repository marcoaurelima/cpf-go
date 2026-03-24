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

	return nil
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

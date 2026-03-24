package cpf

import "errors"

// Declaração de erros específicos para validação de CPF
var (
	ErrorCPFVazio       = errors.New("CPF vazio")
	ErrorCPFIncompleto  = errors.New("CPF com tamanho diferente de 11 dígitos")
	ErrorCPFNaoNumerico = errors.New("CPF contém caracteres não numéricos")
	ErrorCPFInvalido    = errors.New("CPF inválido")
)

// Declaração de erro específico para falha na geração de CPF aleatório
var (
	ErrorGerarCPFAleatorio = errors.New("não foi possível gerar um CPF aleatório válido após o número máximo de tentativas")
)

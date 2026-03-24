package cpf

import "errors"

var (
	ErrorCPFVazio       = errors.New("CPF vazio")
	ErrorCPFIncompleto  = errors.New("CPF com tamanho diferente de 11 dígitos")
	ErrorCPFNaoNumerico = errors.New("CPF contém caracteres não numéricos")
	ErrorCPFInvalido    = errors.New("CPF inválido")
)

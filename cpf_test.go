package cpf

import (
	"testing"
)

func Test_CPF(t *testing.T) {
	t.Run("CPF Válido", func(t *testing.T) {
		CPF, err := New("12345678909")
		if err != nil {
			t.Errorf("Erro ao criar CPF: %v", err)
		}

		if !CPF.IsValid() {
			t.Errorf("CPF deveria ser válido, mas foi considerado inválido")
		}

		if CPF.String() != "12345678909" {
			t.Errorf("String() retornou '%s', esperado '12345678909'", CPF.String())
		}

		if CPF.StringFormatted() != "123.456.789-09" {
			t.Errorf("StringFormatted() retornou '%s', esperado '123.456.789-09'", CPF.StringFormatted())
		}

		if CPF.Base() != "123456789" {
			t.Errorf("Base() retornou '%s', esperado '123456789'", CPF.Base())
		}

		if CPF.DV() != "09" {
			t.Errorf("DV() retornou '%s', esperado '09'", CPF.DV())
		}
	})

	t.Run("CPF com DV inválido", func(t *testing.T) {
		CPF, err := New("12345678908")
		if err == nil {
			t.Errorf("Deveria ter retornado um erro para CPF com DV inválido, mas não retornou")
		}

		if CPF.IsValid() {
			t.Errorf("CPF deveria ser inválido, mas foi considerado válido")
		}
	})

	t.Run("CPF vazio", func(t *testing.T) {
		CPF, err := New("")
		if err == nil {
			t.Errorf("Deveria ter retornado um erro para CPF vazio, mas não retornou")
		}

		if CPF.IsValid() {
			t.Errorf("CPF vazio deveria ser inválido, mas foi considerado válido")
		}
	})

	t.Run("CPF com numeração repetida", func(t *testing.T) {
		CPF, err := New("11111111111")
		if err == nil {
			t.Errorf("Deveria ter retornado um erro para CPF com numeração repetida, mas não retornou")
		}

		if CPF.IsValid() {
			t.Errorf("CPF com numeração repetida deveria ser inválido, mas foi considerado válido")
		}
	})

	t.Run("CPF com numeração incompleta", func(t *testing.T) {
		CPF, err := New("1111111111")
		if err == nil {
			t.Errorf("Deveria ter retornado um erro para CPF com numeração incompleta, mas não retornou")
		}

		if CPF.IsValid() {
			t.Errorf("CPF com numeração incompleta deveria ser inválido, mas foi considerado válido")
		}
	})

	t.Run("CPF com digitos não numéricos", func(t *testing.T) {
		CPF, err := New("1111111111a")
		if err == nil {
			t.Errorf("Deveria ter retornado um erro para CPF com digitos não numéricos, mas não retornou")
		}

		if CPF.IsValid() {
			t.Errorf("CPF com digitos não numéricos deveria ser inválido, mas foi considerado válido")
		}
	})

	t.Run("Geração de CPF aleatório", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			CPF, err := NewRandom()
			if err != nil {
				t.Errorf("Erro ao gerar CPF aleatório: %v", err)
			}
			if !CPF.IsValid() {
				t.Errorf("CPF gerado aleatoriamente deveria ser válido, mas foi considerado inválido")
			}
		}
	})
}

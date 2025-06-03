package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(1, 2)

	if total != 3 {
		t.Errorf("Resultado da soma é inválido: %d", total)
	}
}

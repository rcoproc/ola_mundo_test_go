package main

import "testing"

func TestOla(t *testing.T) {
    resultado := Ola("Ricardo")
    esperado := "Olá, Ricardo"

    if resultado != esperado {
        t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
    }
}

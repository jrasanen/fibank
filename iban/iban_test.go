package iban

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNordeaIban(t *testing.T) {
	iban := Iban{
		Number: "FI2112345600000785",
	}
	iban.MachineReadable()

	assert.Equal(t, "Nordea", iban.Bank.Name)
	assert.Equal(t, "12345600000785151821", iban.MachineFormat)
}

func TestForeignIban(t *testing.T) {
	iban := Iban{
		Number: "NL39RABO0300065264",
	}
	iban.MachineReadable()
	assert.Nil(t, iban.Bank, "Bank should not be found")
	assert.Equal(t, "271011240300065264232139", iban.MachineFormat)
}

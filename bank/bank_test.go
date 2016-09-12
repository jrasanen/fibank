package bank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindBankByID(t *testing.T) {
	bank := FindBankByID(799)

	assert.Equal(t, "Holvi", bank.Name)
}

func TestFindNonExistantID(t *testing.T) {
	bank := FindBankByID(999)

	assert.Nil(t, bank, "Bank should not be found")
}

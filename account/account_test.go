package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCorrectChecksum(t *testing.T) {
	assert.Equal(t, true, validateChecksum("12345600000785"))
}

func TestValidateInvalidChecksum(t *testing.T) {
	assert.Equal(t, false, validateChecksum("x0"))
}

func TestInvalidAccountNumber(t *testing.T) {
	a := Account{
		Number: "12345",
	}

	_, err := a.MachineReadable()
	// If error is not raised, raise an error
	if err == nil {
		t.Error("Error should have been raised")
	}
}

func TestValidAccountNumber(t *testing.T) {
	a := Account{
		Number: "123456-785",
	}

	res, err := a.MachineReadable()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "12345600000785", res)

	b := Account{
		Number: "423456-781",
	}

	res, err = b.MachineReadable()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "42345670000081", res)
}

package iban

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jrasanen/fibank/bank"
)

// http://www.finanssiala.fi/maksujenvalitys/dokumentit/Suomalaiset_rahalaitostunnukset_ja_BIC-koodit.pdf#search=tunnukset

// Letters used in IBAN
var charNumbers = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Helper function to convert IBAN letters to machine readable numbers
// http://www.finanssiala.fi/maksujenvalitys/dokumentit/IBAN_ja_BIC_maksuliikenteessa.pdf
func charToNumber(char string) (int, error) {
	/*
	  A = 10 G = 16 M = 22 S = 28 Y = 34
	  B = 11 H = 17 N = 23 T = 29 Z = 35
	  C = 12 I = 18 O = 24 U = 30
	  D = 13 J = 19 P = 25 V = 31
	  E = 14 K = 20 Q = 26 W = 32
	  F = 15 L = 21 R = 27 X = 33
	*/
	index := strings.Index(charNumbers, strings.ToUpper(char))
	if index == -1 {
		return -1, errors.New("Number not found")
	}
	return index + 10, nil
}

type Iban struct {
	Number        string
	Bank          *bank.Bank
	MachineFormat string // Machine readable representation
}

// MachineReadable returns machine readable representation of IBAN
func (ib *Iban) MachineReadable() (string, error) {
	// Remove spaces
	ibanNumber := strings.Replace(ib.Number, " ", "", -1)

	// Country code is the 2 first characters
	country := ibanNumber[0:2]

	// Country checksum is characters 3-5
	countryChecksum := ibanNumber[2:4]

	// Move country code & country checksum end of iban
	ibanNumber = fmt.Sprintf("%s%s%s", ibanNumber[4:], country, countryChecksum)

	// Checksums & bic codes
	// http://www.finanssiala.fi/maksujenvalitys/dokumentit/Suomalaiset_rahalaitostunnukset_ja_BIC-koodit.pdf#search=tunnukset
	var checksumLength int
	if strings.HasPrefix(ibanNumber, "3") {
		checksumLength = 2
	} else if strings.HasPrefix(ibanNumber, "4") || strings.HasPrefix(ibanNumber, "7") {
		checksumLength = 3
	} else {
		checksumLength = 1
	}

	bankChecksum, _ := strconv.Atoi(ibanNumber[0:checksumLength])

	// Find bank by checksum id
	ib.Bank = bank.FindBankByID(bankChecksum)

	var machineFormat string

	// All letters in IBAN must be converted to numbers
	for _, r := range ibanNumber {
		// Rune to string
		char := string(r)
		toAppend, err := charToNumber(char)
		if err == nil {
			machineFormat = fmt.Sprintf("%s%v", machineFormat, toAppend)
		} else {
			machineFormat = fmt.Sprintf("%s%v", machineFormat, char)
		}
	}

	ib.MachineFormat = machineFormat
	return machineFormat, nil
}

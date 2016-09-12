package main

import (
	"fmt"

	"github.com/jrasanen/fibank/account"
	"github.com/jrasanen/fibank/iban"
)

func main() {
	fmt.Println("Accounts:")
	var accounts []string
	accounts = append(accounts, "423456-781")
	accounts = append(accounts, "123456-785")

	// Convert old account numbers to machine readable form and find out
	// the bank name
	for _, a := range accounts {
		account := account.Account{Bank: nil, Number: a}
		mf, _ := account.MachineReadable()
		fmt.Println(mf + " " + account.Bank.Name)
	}

	fmt.Println("IBAN:")
	var ibans []string
	ibans = append(ibans, "FI2112345600000785")
	ibans = append(ibans, "NL39RABO0300065264")

	// Test convert ibans to machine readable format and find out
	// the bank name
	for _, a := range ibans {
		iban := iban.Iban{Number: a}
		mf, _ := iban.MachineReadable()
		if iban.Bank != nil {
			fmt.Printf("%s %s\n", mf, iban.Bank.Name)
		} else {
			fmt.Printf("%s Foreign bank\n", mf)
		}
	}
}

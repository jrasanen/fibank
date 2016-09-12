package account

// http://www.finanssiala.fi/maksujenvalitys/dokumentit/IBAN_ja_BIC_maksuliikenteessa.pdf

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/jrasanen/fibank/bank"
	"github.com/jrasanen/fibank/modules/pad"
)

var oldAccountFormat *regexp.Regexp

func init() {
	oldAccountFormat, _ = regexp.Compile("([0-9]{6})-([0-9]{2,8})")
}

func validateChecksum(mf string) bool {
	if len(mf) != 14 {
		return false
	}
	weights := []int{2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2}
	//mf_checksum := mf[12:13]

	var prod int

	result := ""

	for i, w := range weights {
		num, _ := strconv.Atoi(string(mf[i]))
		res := strconv.Itoa(num * w)
		result += res
	}

	for _, r := range result {
		num, _ := strconv.Atoi(string(r))
		prod += num
	}

	lastDigit, _ := strconv.Atoi(string(mf[len(mf)-1]))
	if 40-prod == lastDigit {
		return true
	}
	return false
}

func (a *Account) MachineReadable() (string, error) {
	var account = a.Number
	if oldAccountFormat.MatchString(account) == false {
		return "", errors.New("Invalid account number format")
	}

	matches := oldAccountFormat.FindStringSubmatch(account)
	account = fmt.Sprintf("%s%s", matches[1], matches[2])

	firstChecksum, err := strconv.Atoi(string(account[0]))
	if err != nil {
		return "", err
	}

	twoChecksum, err := strconv.Atoi(string(account[0:2]))
	if err != nil {
		return "", err
	}

	var yourBank bank.Bank

	for _, aBank := range bank.FIBanks {
		if aBank.ID == firstChecksum || aBank.ID == twoChecksum {
			yourBank = aBank
			break
		}
	}

	padded := pad.Right(account[:yourBank.Pad], 14-len(account), "0")
	machineFormat := fmt.Sprintf("%s%s", padded, account[yourBank.Pad:])
	a.Bank = &yourBank
	if validateChecksum(machineFormat) == false {
		return "", errors.New("Wrong account checksum")
	}
	return machineFormat, nil
}

// Account
type Account struct {
	Bank   *bank.Bank
	Number string
}

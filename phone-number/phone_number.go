package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Normalize: remove all runes that aren't number
func normalize(phoneNumber string) string {
	re, _ := regexp.Compile(`[0-9]+`)
	return strings.Join(re.FindAllString(phoneNumber, -1), "")
}

// Number: Gives a valid number
func Number(phoneNumber string) (string, error) {
	phoneNumber = normalize(phoneNumber)
	lenght := len(phoneNumber)
	if phoneNumber[0] == '1' && lenght == 11 && phoneNumber[1] > '1' && phoneNumber[4] > '1' || phoneNumber[0] > '1' && lenght == 10 && phoneNumber[3] > '1' {
		return phoneNumber[lenght-10:], nil
	}
	return "", errors.New("invalid phone number")
}

// AreaCode: Gives only area code of a number
func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return phoneNumber[:3], nil
}

// Fomart: return a phone number formated
func Format(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", phoneNumber[0:3], phoneNumber[3:6], phoneNumber[6:]), nil
}

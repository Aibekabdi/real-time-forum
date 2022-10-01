package service

import (
	"errors"
	"net/mail"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func isEmailValid(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return err
	}
	return nil
}

func isPasswordValid(password string) error {
	if len(password) < 8 {
		return errors.New("invalid password")
	}
next:
	for name, classes := range map[string][]*unicode.RangeTable{
		"upper case": {unicode.Upper, unicode.Title},
		"lower case": {unicode.Lower},
		"numeric":    {unicode.Number, unicode.Digit},
	} {
		for _, r := range password {
			if unicode.IsOneOf(classes, r) {
				continue next
			}
		}
		return errors.New("password must have at least one" + name + "character")
	}
	return nil
}

func isEmpty(input string) bool {
	for _, v := range input {
		if !(v <= 32) {
			return false
		}
	}
	return true
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

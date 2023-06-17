package utils

import (
	"errors"
	"net/mail"
	"regexp"
	"unicode"

	"forum/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func IsValidRegister(user *models.User) error {
	if err := isValidEmail(user); err != nil {
		return err
	} else if err := isValidUser(user); err != nil {
		return err
	} else if err := isValidPassword(user); err != nil {
		return err
	}
	var err error
	if user.Password, err = GenerateHashPassword(user.Password); err != nil {
		return err
	}
	return nil
}

func GenerateHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CompareHashAndPassword(hash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return err
	}
	return nil
}

func isValidEmail(user *models.User) error {
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return err
	}
	return nil
}

func isValidUser(user *models.User) error {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", user.Nickname); !ok {
		return errors.New("invalid nickname")
	} else if user.FirstName == "" {
		return errors.New("invalid firstname")
	} else if user.LastName == "" {
		return errors.New("invalid lastname")
	} else if user.Gender != "male" && user.Gender != "female" {
		return errors.New("invalid gender")
	} else if user.Age < 0 || user.Age >= 150 {
		return errors.New("invalid age")
	}
	return nil
}

func isValidPassword(user *models.User) error {
	if len(user.Password) < 8 {
		return errors.New("invalid password")
	}
next:
	for name, classes := range map[string][]*unicode.RangeTable{
		"upper case": {unicode.Upper, unicode.Title},
		"lower case": {unicode.Lower},
		"numeric":    {unicode.Number, unicode.Digit},
	} {
		for _, r := range user.Password {
			if unicode.IsOneOf(classes, r) {
				continue next
			}
		}
		return errors.New("password must have at least one" + name + "character")
	}
	return nil
}

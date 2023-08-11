package validator

import (
	"clean-arch-hex/internal/domain/exception"
	"errors"
	"regexp"
)

func IsEmailValid(e string) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(e) {
		return exception.ErrInvalidEmail
	}
	return nil
}

func CheckValidPassword(password string) error {
	if len(password) < 6 {
		return errors.New("password should contain at least 6 letters")
	}
	if !regexp.MustCompile("[a-z]").MatchString(password) {
		return errors.New("password should contain at least 1 lower case character")
	}
	if !regexp.MustCompile("[A-Z]").MatchString(password) {
		return errors.New("password should contain at least 1 upper case character")
	}
	if !regexp.MustCompile("[0-9]").MatchString(password) {
		return errors.New("password should contain at least 1 number [0-9]")
	}
	if !regexp.MustCompile("[//#$(&}?!{;@)*%]").MatchString(password) {
		return errors.New("password should contain at least 1 special character")
	}
	return nil
}

func CheckValidPhoneNumber(phone string) error {
	patern := `(^01|^05|^07|^21|^25|^27)([0-9]{8})$`
	//`(\+|00)225(01|05|07|21|25|27)([0-9]{8})`
	if !regexp.MustCompile(patern).
		MatchString(phone) {
		return exception.ErrInvalidPhone
	}
	return nil
}

package validation

import "net/mail"

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidPassword(password string) bool {
	return len(password) >= 8 && len(password) <= 32
}

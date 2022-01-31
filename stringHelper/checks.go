package stringHelper

import "net/mail"

// returns true if the given string is an email
func checkEmail(value string) bool {
	_, err := mail.ParseAddress(value)
	return err == nil
}

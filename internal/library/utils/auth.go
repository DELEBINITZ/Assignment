package utils

import (
	"strings"
)

const (
	AdminEmail = "admin@yahoo.com"
)

var AdminEmails = []string{AdminEmail}

func IsAdmin(email string) bool {
	email = strings.ToLower(email)

	for _, adminEmail := range AdminEmails {
		if email == adminEmail {
			return true
		}
	}
	return false
}

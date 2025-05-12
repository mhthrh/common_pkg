package validation

import (
	"regexp"
	"strings"
)

var (
	emailRegex  = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	mobileRegex = regexp.MustCompile(`^1?[2-9][0-9]{2}[2-9][0-9]{6}$`)
	nameRegex   = regexp.MustCompile(`^[A-Za-z][A-Za-z'\- ]*[A-Za-z]$`)
	uppercase   = regexp.MustCompile(`[A-Z]`)
	lowercase   = regexp.MustCompile(`[a-z]`)
	digit       = regexp.MustCompile(`[0-9]`)
	specialChar = regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]{};':"\\|,.<>\/?]`)
)

func MobilePhone(number string) bool {
	cleaned := strings.ReplaceAll(number, " ", "")
	cleaned = strings.ReplaceAll(cleaned, "-", "")
	cleaned = strings.ReplaceAll(cleaned, "(", "")
	cleaned = strings.ReplaceAll(cleaned, ")", "")

	return mobileRegex.MatchString(cleaned)
}

func Name(name string) bool {
	name = strings.TrimSpace(name)

	if len(name) < 2 || len(name) > 50 {
		return false
	}

	return nameRegex.MatchString(name)
}

func Email(email string) bool {
	email = strings.TrimSpace(email)
	return emailRegex.MatchString(email)
}
func Password(password string) bool {
	if len(password) < 8 {
		return false
	}

	return uppercase.MatchString(password) &&
		lowercase.MatchString(password) &&
		digit.MatchString(password) &&
		specialChar.MatchString(password)
}

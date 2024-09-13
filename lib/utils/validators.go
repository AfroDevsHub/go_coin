package utils

import (
	"fmt"

	"github.com/dlclark/regexp2"
)

func ValidateEmail(email string) error {
	re := regexp2.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-_]+\.[a-zA-Z]{1,3}\.?[a-zA-Z]{2,3}$`, regexp2.IgnoreCase)
	if isMatch, _ := re.MatchString(email); !isMatch {
		return fmt.Errorf("invalid email")
	}
	return nil
}

func ValidatePassword(password string) error {
	re := regexp2.MustCompile(`^(?=.*[a-zA-Z])(?=.*\d)(?=.*[!@#$%^&*()-_+=]).{8,}$`, regexp2.None)
	if isMatch, _ := re.MatchString(password); !isMatch {
		return fmt.Errorf("invalid password")
	}
	return nil
}

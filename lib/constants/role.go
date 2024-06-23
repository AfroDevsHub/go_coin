package constants

import "fmt"

type Role int

const (
	ADMIN Role = iota
	USER
	DEVELOPER
)

var roles = map[Role]string{
	ADMIN:     "Admin User.",
	USER:      "User.",
	DEVELOPER: "Developer.",
}

func (r Role) String() (string, error) {
	if value, ok := roles[r]; ok {
		return value, nil
	}
	return "", fmt.Errorf("invalid role")
}

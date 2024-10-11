package constants

import "fmt"

type Role int

const (
	USER Role = iota
	DEVELOPER
	ADMIN
)

var roles = map[Role]string{
	ADMIN:     "Admin",
	USER:      "User",
	DEVELOPER: "Developer",
}

func (r Role) String() (string, error) {
	if value, ok := roles[r]; ok {
		return value, nil
	}
	return "", fmt.Errorf("invalid role")
}

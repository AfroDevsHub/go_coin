package constants

import "fmt"

type ProfileVisibility int

const (
	PRIVATE ProfileVisibility = iota
	PUBLIC
	PROFILEVISIBILITYADMIN
)

var profilevisibilities = map[ProfileVisibility]string{
	PRIVATE:                "Private",
	PUBLIC:                 "Public",
	PROFILEVISIBILITYADMIN: "Admin",
}

func (p ProfileVisibility) String() (string, error) {
	if value, ok := profilevisibilities[p]; ok {
		return value, nil
	}
	return "", fmt.Errorf("invalid profile visibility")
}

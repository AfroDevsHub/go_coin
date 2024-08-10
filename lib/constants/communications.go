package constants

import "fmt"

type Communication int

const (
	COMMUNICATIONEMAIL Communication = iota
	SMS
	CELL
	COMMUNICATIONSLACK
)

var communications = map[Communication]string{
	COMMUNICATIONEMAIL: "Email Messenger",
	SMS:                "SMS",
	CELL:               "Mobile Phone",
	COMMUNICATIONSLACK: "Slack Messenger",
}

func (c Communication) String() (string, error) {
	if value, ok := communications[c]; ok {
		return value, nil
	}
	return "", fmt.Errorf("invalid communication")
}

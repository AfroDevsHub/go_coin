package constants

import "fmt"

type LoginMethod int

const (
	LOGINEMAIL LoginMethod = iota
	GITHUB
	LOGINSLACK
	GOOGLE
	FACEBOOK
)

var loginmethods = map[LoginMethod]string{
	LOGINEMAIL: "User Email and Password",
	GITHUB:     "Github SSO",
	LOGINSLACK: "Slack SSO",
	GOOGLE:     "Google SSO",
	FACEBOOK:   "Facebook SSO",
}

func (l LoginMethod) String() (string, error) {
	if value, ok := loginmethods[l]; ok {
		return value, nil
	}
	return "", fmt.Errorf("invalid login method")
}

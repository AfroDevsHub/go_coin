package constants

import "fmt"

type DataSharingPreference int

const (
	ACCOUNT DataSharingPreference = iota
	PROFILE
	SETTINGS
	TRANSACTIONS
)

var datasharingpreferences = map[DataSharingPreference]string{
	ACCOUNT:      "Share All Permissible Account Data.",
	PROFILE:      "Share Profile Data.",
	SETTINGS:     "Share Settings Data.",
	TRANSACTIONS: "Share Transaction Data.",
}

func (d DataSharingPreference) String() (string, error) {
	if value, ok := datasharingpreferences[d]; ok {
		return value, nil
	}
	return "", fmt.Errorf("invalid data sharing preference")
}

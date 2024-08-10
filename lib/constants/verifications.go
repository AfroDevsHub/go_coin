package constants

import "fmt"

type Verification int

const (
	UNVERIFIED Verification = iota
	VERIFIED
	VERIFYING
	FAILED
	EXPIRED
)

var verifications = map[Verification]string{
	UNVERIFIED: "New and Unverified",
	VERIFIED:   "Verified",
	VERIFYING:  "Verification Requested",
	FAILED:     "Verification Failed",
	EXPIRED:    "Verification Request Expired",
}

func (v Verification) String() (string, error) {
	if value, ok := verifications[v]; ok {
		return value, nil
	}
	return "", fmt.Errorf("invalid verification")
}

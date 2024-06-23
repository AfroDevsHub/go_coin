package constants

import (
	"fmt"
)

type Status int

const (
	NEW Status = iota
	ACTIVE
	INACTIVE
	DISABLED
	DELETED
)

var statuses = map[Status]string{
	NEW:      "Newly Created.",
	ACTIVE:   "Actively is in Use.",
	INACTIVE: "Not Actively in Use.",
	DISABLED: "Can't be used.",
	DELETED:  "Has been Deleted.",
}

func (s Status) String() (string, error) {
	if value, ok := statuses[s]; ok {
		return value, nil
	}
	return "", fmt.Errorf("invalid status")
}

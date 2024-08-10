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
	NEW:      "New",
	ACTIVE:   "Active",
	INACTIVE: "Inactive",
	DISABLED: "Disabled",
	DELETED:  "Deleted",
}

func (s Status) String() (string, error) {
	if value, ok := statuses[s]; ok {
		return value, nil
	}
	return "", fmt.Errorf("invalid status")
}

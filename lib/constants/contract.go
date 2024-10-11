package constants

import "fmt"

type ContractType int

const (
	CONTRACT_DRAFT ContractType = iota
	CONTRACT_APPROVED
	CONTRACT_REJECTED
	CONTRACT_ACTIVE
	CONTRACT_CLOSED
)

var contract_type = map[ContractType]string{
	CONTRACT_DRAFT:    "Draft",
	CONTRACT_APPROVED: "Approved",
	CONTRACT_REJECTED: "Rejected",
	CONTRACT_ACTIVE:   "Active",
	CONTRACT_CLOSED:   "Closed",
}

func (c ContractType) String() (string, error) {
	if value, ok := contract_type[c]; ok {
		return value, nil
	}
	return "", fmt.Errorf("invalid contract type")
}

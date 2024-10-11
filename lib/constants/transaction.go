package constants

import "fmt"

type TransactionType int

const (
	TRANSACTION_DRAFT TransactionType = iota
	TRANSACTION_APPROVED
	TRANSACTION_REJECTED
	TRANSACTION_INSUFFICIENT
	TRANSACTION_TRANSFERRED
	TRANSACTION_REVERSED
)

var transaction_type = map[TransactionType]string{
	TRANSACTION_DRAFT:        "Transaction Drafted.",
	TRANSACTION_APPROVED:     "Transaction Approved.",
	TRANSACTION_REJECTED:     "Transaction Rejected.",
	TRANSACTION_INSUFFICIENT: "Insufficient Funds.",
	TRANSACTION_TRANSFERRED:  "Funds Transferred.",
	TRANSACTION_REVERSED:     "Transaction Reversed.",
}

func (c TransactionType) String() (string, error) {
	if value, ok := transaction_type[c]; ok {
		return value, nil
	}
	return "", fmt.Errorf("invalid transaction type")
}

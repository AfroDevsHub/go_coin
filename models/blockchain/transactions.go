package blockchain

import (
	"time"

	"github.com/google/uuid"
)

type transaction struct {
	id                 uuid.UUID
	transaction_id     uuid.UUID
	sender             uuid.UUID
	receiver           uuid.UUID
	amount             float64
	title              string
	description        string
	sender_signiture   string
	receiver_signiture string
	transaction_status string
	salt_value         uuid.UUID
	created_date       time.Time
	updated_date       time.Time
}

func (t *transaction) String() string {
	return "Transaction ID: " + t.transaction_id.String()
}

func (t *transaction) Dict() map[string]interface{} {
	return map[string]interface{}{
		"id":                 t.id,
		"transaction_id":     t.transaction_id,
		"sender":             t.sender,
		"receiver":           t.receiver,
		"amount":             t.amount,
		"title":              t.title,
		"description":        t.description,
		"sender_signiture":   t.sender_signiture,
		"receiver_signiture": t.receiver_signiture,
		"transaction_status": t.transaction_status,
		"salt_value":         t.salt_value,
		"created_date":       t.created_date,
		"updated_date":       t.updated_date,
	}
}

type Transaction struct{}

func (t *Transaction) create_transaction() (*transaction, error) {
	return &transaction{
		uuid.New(),
		uuid.New(),
		uuid.New(),
		uuid.New(),
		1.02,
		"String",
		"String",
		"String",
		"String",
		"String",
		uuid.New(),
		time.Now(),
		time.Now(),
	}, nil
}

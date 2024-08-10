package blockchain

import (
	"time"

	"github.com/google/uuid"
)

type contract struct {
	id                   uuid.UUID
	contract_id          uuid.UUID
	contractor           uuid.UUID
	contractee           uuid.UUID
	amount               float64
	title                string
	description          string
	contractor_signiture string
	contractee_signiture string
	contract_status      string
	salt_value           uuid.UUID
	created_date         time.Time
	updated_date         time.Time
}

func (c *contract) String() string {
	return "Contract ID: " + c.contract_id.String()
}

func (c *contract) Dict() map[string]interface{} {
	return map[string]interface{}{
		"id":                   c.id,
		"contract_id":          c.contract_id,
		"contractor":           c.contractor,
		"contractee":           c.contractee,
		"amount":               c.amount,
		"title":                c.title,
		"description":          c.description,
		"contractor_signiture": c.contractor_signiture,
		"contractee_signiture": c.contractee_signiture,
		"contract_status":      c.contract_status,
		"salt_value":           c.salt_value,
		"created_date":         c.created_date,
		"updated_date":         c.updated_date,
	}
}

type Contract struct{}

func (c *Contract) create_contract() (*contract, error) {
	return &contract{
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

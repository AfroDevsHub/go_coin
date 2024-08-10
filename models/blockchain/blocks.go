package blockchain

import (
	"time"

	"github.com/google/uuid"
)

type block struct {
	id                uuid.UUID
	block_id          uuid.UUID
	transaction_id    uuid.UUID
	contract_id       uuid.UUID
	previous_block_id uuid.UUID
	next_block_id     uuid.UUID
	block_type        string
	created_date      time.Time
	updated_date      time.Time
}

func (b *block) String() string {
	return "Block ID: " + b.block_id.String()
}

func (b *block) Dict() map[string]interface{} {
	return map[string]interface{}{
		"id":                b.id,
		"block_id":          b.block_id,
		"transaction_id":    b.transaction_id,
		"contract_id":       b.contract_id,
		"previous_block_id": b.previous_block_id,
		"next_block_id":     b.next_block_id,
		"block_type":        b.block_type,
		"created_date":      b.created_date,
		"updated_date":      b.updated_date,
	}
}

type Block struct{}

func (b *Block) create_block() (*block, error) {
	return &block{
		uuid.New(),
		uuid.New(),
		uuid.New(),
		uuid.New(),
		uuid.New(),
		uuid.New(),
		"String",
		time.Now(),
		time.Now(),
	}, nil
}

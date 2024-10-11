package blockchain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Block struct {
	ID              uuid.UUID    `json:"id" gorm:"type:uuid;not null"`
	BlockID         uuid.UUID    `json:"block_id" gorm:"type:uuid;not null"`
	ContractID      *uuid.UUID   `json:"contract_id" gorm:"type:uuid"`
	TransactionID   *uuid.UUID   `json:"transaction_id" gorm:"type:uuid"`
	PreviousBlockID *uuid.UUID   `json:"previous_block_id" gorm:"type:uuid"`
	NextBlockID     *uuid.UUID   `json:"next_block_id" gorm:"type:uuid"`
	BlockType       string       `json:"block_type" gorm:"type:string;not null"`
	SaltValue       uuid.UUID    `json:"salt_value" gorm:"type:uuid;not null"`
	Contract        *Contract    `gorm:"references:ID"`
	Transaction     *Transaction `gorm:"references:ID"`
	gorm.Model
}

func (*Block) TableName() string {
	return "blockchain.block"
}

func (b *Block) String() string {
	return "Block ID: " + b.BlockID.String()
}

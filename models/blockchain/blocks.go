package blockchain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Block struct {
	ID              uuid.UUID    `json:"id" gorm:"type=uuid;not null"`
	BlockID         uuid.UUID    `json:"block_id" gorm:"type=uuid;not null"`
	ContractID      *uuid.UUID   `json:"contract_id"`
	TransactionID   *uuid.UUID   `json:"transaction_id"`
	PreviousBlockID *uuid.UUID   `json:"previousBlock_id"`
	NextBlockID     *uuid.UUID   `json:"nextBlock_id"`
	BlockType       string       `json:"block_type" gorm:"type:string;not null"`
	SaltValue       uuid.UUID    `json:"salt_value" gorm:"type=uuid;not null"`
	Transaction     *Transaction `gorm:"foreignKey:TransactionID;references:ID"`
	Contract        *Contract    `gorm:"foreignKey:ContractID;references:ID"`
	gorm.Model
}

func (*Block) TableName() string {
	return "blockchain.contracts"
}

func (b *Block) String() string {
	return "Block ID: " + b.BlockID.String()
}

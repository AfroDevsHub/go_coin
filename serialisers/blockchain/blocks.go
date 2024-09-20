package blockchain

import (
	"fmt"

	"github.com/dfunani/go_coin/database"
	"github.com/dfunani/go_coin/lib/constants"
	"github.com/dfunani/go_coin/models/blockchain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlockSerialiser struct{}

func (*BlockSerialiser) ReadBlock(db *gorm.DB, block_id uuid.UUID) (*blockchain.Block, error) {
	var block []blockchain.Block
	db.Model(&blockchain.Block{}).Find(&block, &blockchain.Block{BlockID: block_id})
	if len(block) != 1 {
		return &blockchain.Block{}, fmt.Errorf("invalid block id")
	}
	return &block[0], nil
}

func (b *BlockSerialiser) CreateBlock(db *gorm.DB, transaction_id *uuid.UUID, contract_id *uuid.UUID, PreviousBlockID *uuid.UUID,
	NextBlockID *uuid.UUID) (string, error) {
	var block_type string
	if transaction_id == nil && contract_id == nil {
		block_type, _ = constants.UNIT.String()
	} else if transaction_id != nil && contract_id != nil {
		return "", fmt.Errorf("invalid block")
	} else if contract_id != nil {
		block_type, _ = constants.CONTRACT.String()
	} else {
		block_type, _ = constants.TRANSACTION.String()
	}
	return database.CreateModel(db, &blockchain.Block{
		ID:              uuid.New(),
		BlockID:         uuid.New(),
		ContractID:      contract_id,
		TransactionID:   transaction_id,
		PreviousBlockID: PreviousBlockID,
		NextBlockID:     NextBlockID,
		BlockType:       block_type,
		SaltValue:       uuid.New(),
	})
}

func (*BlockSerialiser) UpdateBlock(db *gorm.DB, id uuid.UUID, block_data map[string]interface{}) (string, error) {
	var block blockchain.Block
	err := db.Model(&blockchain.Block{}).First(&block, id).Error
	if err != nil {
		return "", err
	}
	db.Model(block).Updates(block_data)
	return block.String(), nil
}

func (*BlockSerialiser) DeleteBlock(db *gorm.DB, id uuid.UUID) (string, error) {
	err := db.Delete(&blockchain.Block{}, id).Error
	if err != nil {
		return "", err
	}
	return id.String() + " Deleted", nil
}

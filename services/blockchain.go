package services

import (
	"github.com/dfunani/go_coin/lib/constants"
	"github.com/dfunani/go_coin/models/blockchain"
	SERIALISERS "github.com/dfunani/go_coin/serialisers/blockchain"
	"gorm.io/gorm"
)

var blockChain = struct {
	blocks []blockchain.Block
}{
	blocks: []blockchain.Block{},
}

func CreateBlock(db *gorm.DB, block_type constants.BlockType, block interface{}) (string, error) {
	serialiser := SERIALISERS.BlockSerialiser{}
	switch result := block.(type) {
	case blockchain.Contract:
		return serialiser.CreateBlock(db, nil, &result.ID, nil, nil)
	case blockchain.Transaction:
		return serialiser.CreateBlock(db, &result.ID, nil, nil, nil)
	}
	return serialiser.CreateBlock(db, nil, nil, nil, nil)

}

func GetBlockChain() []blockchain.Block {
	return blockChain.blocks
}

func UpdateBlockChain() {
	blockChain.blocks = append(blockChain.blocks, blockchain.Block{})
}

package constants

import "fmt"

type BlockType int

const (
	TRANSACTION BlockType = iota
	CONTRACT
	UNIT
)

var block_type = map[BlockType]string{
	TRANSACTION: "Transaction Block",
	CONTRACT:    "Contract Block",
	UNIT:        "Unit Block",
}

func (c BlockType) String() (string, error) {
	if value, ok := block_type[c]; ok {
		return value, nil
	}
	return "", fmt.Errorf("invalid block type")
}

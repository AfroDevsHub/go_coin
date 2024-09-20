package blockchain

import (
	"fmt"

	"github.com/dfunani/go_coin/database"
	"github.com/dfunani/go_coin/lib/constants"
	"github.com/dfunani/go_coin/models/blockchain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContractSerialiser struct{}

func (*ContractSerialiser) ReadContract(db *gorm.DB, contract_id uuid.UUID) (*blockchain.Contract, error) {
	var contract []blockchain.Contract
	db.Model(&blockchain.Contract{}).Find(&contract, &blockchain.Contract{ContractID: contract_id})
	if len(contract) != 1 {
		return &blockchain.Contract{}, fmt.Errorf("invalid contract id")
	}
	return &contract[0], nil
}

func (c *ContractSerialiser) CreateContract(db *gorm.DB, contractor_id uuid.UUID, contractee_id uuid.UUID, title string, description string, contract_data string) (string, error) {
	var status string
	status, _ = constants.CONTRACT_DRAFT.String()
	return database.CreateModel(db, &blockchain.Contract{
		ID:                  uuid.New(),
		ContractID:          uuid.New(),
		ContractorID:        contractor_id,
		ContracteeID:        contractee_id,
		Title:               title,
		Description:         description,
		ContractorSigniture: "String",
		ContracteeSigniture: "String",
		ContractStatus:      status,
		ContractData:        contract_data,
		Salt_value:          uuid.New(),
	})
}

func (*ContractSerialiser) UpdateContract(db *gorm.DB, id uuid.UUID, contract_data map[string]interface{}) (string, error) {
	var contract blockchain.Contract
	err := db.Model(&blockchain.Contract{}).First(&contract, id).Error
	if err != nil {
		return "", err
	}
	db.Model(contract).Updates(contract_data)
	return contract.String(), nil
}

func (*ContractSerialiser) DeleteContract(db *gorm.DB, id uuid.UUID) (string, error) {
	err := db.Delete(&blockchain.Contract{}, id).Error
	if err != nil {
		return "", err
	}
	return id.String() + " Deleted", nil
}

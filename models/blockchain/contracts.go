package blockchain

import (
	users "github.com/dfunani/go_coin/models/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Contract struct {
	ID                  uuid.UUID            `json:"id" gorm:"type=uuid;not null"`
	ContractID          uuid.UUID            `json:"contract_id" gorm:"type=uuid;not null"`
	ContractorID        uuid.UUID            `json:"contractor" gorm:"type=uuid;not null"`
	ContracteeID        uuid.UUID            `json:"contractee" gorm:"type=uuid;not null"`
	Title               string               `json:"title" gorm:"type=string;not null"`
	Description         string               `json:"description" gorm:"type=string;not null"`
	ContractorSigniture string               `json:"contractor_signiture" gorm:"type=string;not null"`
	ContracteeSigniture string               `json:"contractee_signiture" gorm:"type=string;not null"`
	ContractStatus      string               `json:"contract_status" gorm:"type=string;not null"`
	ContractData        string               `json:"amount" gorm:"not null"`
	Salt_value          uuid.UUID            `json:"salt_value" gorm:"type=uuid;not null"`
	Contractor          users.PaymentProfile `gorm:"foreignKey:ContractorID;references:ID"`
	Contractee          users.PaymentProfile `gorm:"foreignKey:ContracteeID;references:ID"`
	gorm.Model
}

func (*Contract) TableName() string {
	return "blockchain.contracts"
}

func (c *Contract) String() string {
	return "Contract ID: " + c.ContractID.String()
}

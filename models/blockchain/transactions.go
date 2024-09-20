package blockchain

import (
	users "github.com/dfunani/go_coin/models/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID                uuid.UUID            `json:"id" gorm:"type=uuid;not null"`
	TransactionID     uuid.UUID            `json:"transaction_id" gorm:"type=uuid;not null"`
	SenderID          uuid.UUID            `json:"sender" gorm:"type=uuid;not null"`
	ReceiverID        uuid.UUID            `json:"receiver" gorm:"type=uuid;not null"`
	Title             string               `json:"title" gorm:"type=string;not null"`
	Description       string               `json:"description" gorm:"type=string;not null"`
	SenderSigniture   string               `json:"sender_signiture" gorm:"type=string;not null"`
	ReceiverSigniture string               `json:"receiver_signiture" gorm:"type=string;not null"`
	TransactionStatus string               `json:"transaction_status" gorm:"type=string;not null"`
	Amount            float32              `json:"amount" gorm:"not null"`
	Salt_value        uuid.UUID            `json:"salt_value" gorm:"type=uuid;not null"`
	Sender            users.PaymentProfile `gorm:"foreignKey:SenderID;references:ID"`
	Receiver          users.PaymentProfile `gorm:"foreignKey:ReceiverID;references:ID"`
	gorm.Model
}

func (*Transaction) TableName() string {
	return "blockchain.transactions"
}

func (t *Transaction) String() string {
	return "Transaction ID: " + t.TransactionID.String()
}

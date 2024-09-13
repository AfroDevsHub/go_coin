package users

import (
	"github.com/dfunani/go_coin/models/warehouse"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentProfile struct {
	ID          uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey"`
	PaymentID   uuid.UUID       `json:"payment_id" gorm:"type:uuid;uniqueIndex;not null"`
	AccountID   uuid.UUID       `json:"account_id" gorm:"type:uuid;not null"`
	CardID      uuid.UUID       `json:"card_id" gorm:"type:uuid;uniqueIndex;not null"`
	Name        string          `json:"name" gorm:"type:string;not null"`
	Description string          `json:"description" gorm:"type:string;not null"`
	Status      string          `json:"status" gorm:"type:string;not null"`
	Balance     float64         `json:"balance" gorm:"not null;default:0.00"`
	Card        *warehouse.Card `gorm:"references:ID"`
	gorm.Model
}

func (l *PaymentProfile) TableName() string {
	return "users.payment_profiles"
}

func (p *PaymentProfile) String() string {
	return "Payment Profile ID: " + p.PaymentID.String()
}

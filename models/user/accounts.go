package users

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID              uuid.UUID        `json:"id" gorm:"type:uuid;primaryKey"`
	AccountID       uuid.UUID        `json:"account_id" gorm:"type:uuid;not null;uniqueIndex"`
	UserID          uuid.UUID        `json:"user_id" gorm:"type:uuid;not null;uniqueIndex"`
	Status          string           `json:"status" gorm:"type:string;not null"`
	UserProfile     UserProfile      `gorm:"foreignKey:AccountID;references:ID"`
	PaymentProfiles []PaymentProfile `gorm:"foreignKey:AccountID;references:ID"`
	SettingsProfile SettingsProfile  `gorm:"foreignKey:AccountID;references:ID"`
	gorm.Model
}

func (*Account) TableName() string {
	return "users.accounts"
}

func (a *Account) String() string {
	return "Account ID: " + a.AccountID.String()
}

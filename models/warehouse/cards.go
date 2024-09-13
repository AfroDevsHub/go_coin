package warehouse

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Card struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	CardID         uuid.UUID `json:"card_id" gorm:"type:uuid;uniqueIndex;not null"`
	CardNumber     string    `json:"card_number" gorm:"type:string;not null"`
	CvvNumber      string    `json:"cvv_number" gorm:"type:string;not null"`
	CardType       string    `json:"card_type" gorm:"type:string;not null"`
	Status         string    `json:"status" gorm:"type:string;not null"`
	Pin            string    `json:"pin" gorm:"type:string;not null"`
	ExpirationDate time.Time `json:"expiration_date" gorm:"type:timestamp;not null"`
	SaltValue      uuid.UUID `json:"salt_value" gorm:"type:uuid;not null"`
	gorm.Model
}

func (c *Card) TableName() string {
	return "warehouse.cards"
}

func (c *Card) String() string {
	return "Card ID: " + c.CardID.String()
}

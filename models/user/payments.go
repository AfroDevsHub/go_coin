package users

import (
	"github.com/dfunani/go_coin/lib/constants"
	"gorm.io/gorm"

	// "github.com/dfunani/go_coin/models/warehouse"
	"github.com/google/uuid"
)

type PaymentProfile struct {
	gorm.Model
	ID          uuid.UUID
	PaymentID   uuid.UUID
	AccountID   uuid.UUID
	CardID      interface{}
	Name        string
	Description string
	Status      constants.Status
	Balance     float32
}

func (l *PaymentProfile) TableName() string {
	return "users.payment_profiles"
}

func (p *PaymentProfile) String() string {
	return "Payment Profile ID: " + p.PaymentID.String()
}

func (p *PaymentProfile) Dict() map[string]interface{} {
	return map[string]interface{}{
		"id":          p.ID,
		"payment_id":  p.PaymentID,
		"account_id":  p.AccountID,
		"card_id":     p.CardID,
		"name":        p.Name,
		"description": p.Description,
		"status":      p.Status,
		"balance":     p.Balance,
	}
}

type PaymentProfileSerialiser struct{}

func (*PaymentProfileSerialiser) Create_payment_profile() (*PaymentProfile, error) {
	// cards := warehouse.Card{}
	// card, err := cards.Create_card(constants.CardTypes.SAVINGS, "256381")
	// if err != nil {
	// 	return &PaymentProfile{}, fmt.Errorf("card error")
	// }
	return &PaymentProfile{
		ID:          uuid.New(),
		PaymentID:   uuid.New(),
		AccountID:   uuid.New(),
		CardID:      "card.Dict()[id]",
		Name:        "Card Holder",
		Description: "Describes the Payment Profile.",
		Status:      constants.Statuses.NEW,
		Balance:     0.0,
	}, nil
}

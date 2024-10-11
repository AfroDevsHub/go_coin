package serialisers

import (
	"fmt"

	"github.com/dfunani/go_coin/database"
	"github.com/dfunani/go_coin/lib/constants"
	"github.com/dfunani/go_coin/lib/utils"
	users "github.com/dfunani/go_coin/models/user"
	"github.com/dfunani/go_coin/models/warehouse"
	WAREHOUSE_SERIALISER "github.com/dfunani/go_coin/serialisers/warehouse"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentProfileSerialiser struct{}

func (*PaymentProfileSerialiser) ReadPaymentProfile(db *gorm.DB, payment_id uuid.UUID) (*users.PaymentProfile, error) {
	var payment []users.PaymentProfile
	db.Model(&users.PaymentProfile{}).Preload("Card").Find(&payment, users.PaymentProfile{PaymentID: payment_id})
	if len(payment) != 1 {
		return &users.PaymentProfile{}, fmt.Errorf("invalid payment id")
	}
	return &payment[0], nil
}

func (*PaymentProfileSerialiser) CreatePaymentProfile(db *gorm.DB, account_id uuid.UUID, card_type constants.CardType, pin string) (string, error) {
	cards := WAREHOUSE_SERIALISER.CardSerialiser{}
	card, err := cards.CreateCard(db, card_type, pin)
	if err != nil {
		return "", fmt.Errorf("card error - %s", err)
	}
	card_id, err := utils.GetModelID(card)
	if err != nil {
		return "", fmt.Errorf("invalid card id")
	}
	var get_card warehouse.Card
	db.Model(&warehouse.Card{}).First(&get_card, warehouse.Card{CardID: uuid.MustParse(card_id)})

	var valid_status string
	if valid_status, err = constants.Statuses.NEW.String(); err != nil {
		return "", err
	}
	return database.CreateModel(db, &users.PaymentProfile{
		ID:          uuid.New(),
		PaymentID:   uuid.New(),
		AccountID:   account_id,
		CardID:      get_card.ID,
		Name:        "Card Holder",
		Description: "Describes the Payment Profile.",
		Status:      valid_status,
		Balance:     0.0,
	})
}

func (*PaymentProfileSerialiser) UpdatePaymentProfile(db *gorm.DB, id uuid.UUID, payment_profile_data map[string]interface{}) (string, error) {
	var payment_profile users.PaymentProfile
	err := db.Model(&users.PaymentProfile{}).First(&payment_profile, id).Error
	if err != nil {
		return "", err
	}
	db.Model(payment_profile).Updates(payment_profile_data)
	return payment_profile.String(), nil
}

func (*PaymentProfileSerialiser) DeletePaymentProfile(db *gorm.DB, id uuid.UUID) (string, error) {
	err := db.Delete(&users.PaymentProfile{}, id).Error
	if err != nil {
		return "", err
	}
	return id.String() + " Deleted", nil
}

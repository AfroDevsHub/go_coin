package warehouse

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/dfunani/go_coin/lib/constants"
	"github.com/dlclark/regexp2"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

var card_length = 9
var ccv_length = 3

type Card struct {
	gorm.Model
	ID             uuid.UUID
	CardID         uuid.UUID
	CardNumber     string
	CvvNumber      string
	CardType       constants.CardType
	Status         constants.Status
	Pin            string
	ExpirationDate time.Time
	SaltValue      uuid.UUID
}

func (c *Card) TableName() string {
	return "warehouse.cards"
}

func (c *Card) String() string {
	return "Card ID: " + c.CardID.String()
}

func (c *Card) Dict() map[string]interface{} {
	return map[string]interface{}{
		"id":              c.ID,
		"card_id":         c.CardID,
		"card_number":     c.CardNumber,
		"cvv_number":      c.CvvNumber,
		"card_type":       c.CardType,
		"status":          c.Status,
		"pin":             c.Pin,
		"expiration_date": c.ExpirationDate.Format("06/01"),
		"salt_value":      c.SaltValue,
	}
}

type CardSerialiser struct{}

func (*CardSerialiser) Create_card(card_type constants.CardType, pin string) (*Card, error) {
	card_number, err := generate_card_number(card_type)
	if err != nil {
		return &Card{}, fmt.Errorf("invalid card")
	}
	re := regexp2.MustCompile(`^\d{6}$`, regexp2.None)
	if isMatch, _ := re.MatchString(pin); !isMatch {
		return &Card{}, fmt.Errorf("invalid pin")
	}
	return &Card{
		ID:             uuid.New(),
		CardID:         uuid.New(),
		CardNumber:     card_number,
		CvvNumber:      generate_cvv_number(),
		CardType:       card_type,
		Status:         constants.Statuses.ACTIVE,
		Pin:            pin,
		ExpirationDate: time.Now().AddDate(5, 0, 0),
		SaltValue:      uuid.New(),
	}, nil
}

func generate_card_number(card_type constants.CardType) (string, error) {
	value, err := card_type.String()
	if err != nil {
		return "", fmt.Errorf("no card number generated")
	}
	generic := ""
	length := make([]int, card_length)
	for range length {
		generic += strconv.Itoa(rand.Intn(9))

	}
	return value + generic, nil
}

func generate_cvv_number() string {
	generic := ""
	length := make([]int, ccv_length)
	for range length {
		generic += strconv.Itoa(rand.Intn(9))
	}
	return generic
}

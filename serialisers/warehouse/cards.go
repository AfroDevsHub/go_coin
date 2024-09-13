package serialisers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dfunani/go_coin/database"
	"github.com/dfunani/go_coin/lib/constants"
	"github.com/dfunani/go_coin/models/warehouse"
	"github.com/dlclark/regexp2"
	"github.com/google/uuid"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

type CardSerialiser struct{}

func (*CardSerialiser) ReadCard(db *gorm.DB, card_id uuid.UUID) (*warehouse.Card, error) {
	var card []warehouse.Card
	db.Model(&warehouse.Card{}).Find(&card, &warehouse.Card{CardID: card_id})
	if len(card) != 1 {
		return &warehouse.Card{}, fmt.Errorf("invalid card id")
	}
	return &card[0], nil
}

func (*CardSerialiser) CreateCard(db *gorm.DB, card_type constants.CardType, pin string) (string, error) {
	card_number, err := generate_card_number(card_type)
	if err != nil {
		return "", fmt.Errorf("invalid card")
	}
	re := regexp2.MustCompile(`^\d{6}$`, regexp2.None)
	if isMatch, _ := re.MatchString(pin); !isMatch {
		return "", fmt.Errorf("invalid pin")
	}

	var valid_card_type string
	if valid_card_type, err = card_type.String(); err != nil {
		return "", err
	}

	var valid_card_status string
	if valid_card_status, err = constants.Statuses.ACTIVE.String(); err != nil {
		return "", err
	}
	return database.CreateModel(db, &warehouse.Card{
		ID:             uuid.New(),
		CardID:         uuid.New(),
		CardNumber:     card_number,
		CvvNumber:      generate_cvv_number(),
		CardType:       valid_card_type,
		Status:         valid_card_status,
		Pin:            pin,
		ExpirationDate: time.Now().AddDate(5, 0, 0),
		SaltValue:      uuid.New(),
	})
}

func generate_card_number(card_type constants.CardType) (string, error) {
	value, err := card_type.String()
	if err != nil {
		return "", fmt.Errorf("no card number generated")
	}
	generic := ""
	length := make([]int, constants.CARD_LENGTH)
	for range length {
		generic += strconv.Itoa(rand.Intn(9))

	}
	return value + generic, nil
}

func generate_cvv_number() string {
	generic := ""
	length := make([]int, constants.CCV_LENGTH)
	for range length {
		generic += strconv.Itoa(rand.Intn(9))
	}
	return generic
}

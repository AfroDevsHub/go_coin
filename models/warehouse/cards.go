package warehouse

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/dfunani/go_coin/lib/constants"
	"github.com/dlclark/regexp2"
	"github.com/google/uuid"
)

type card struct {
	id              uuid.UUID
	card_id         uuid.UUID
	card_number     string
	cvv_number      string
	card_type       constants.CardType
	status          constants.Status
	pin             string
	expiration_date time.Time
	salt_value      uuid.UUID
	created_date    time.Time
	updated_date    time.Time
}

func (c *card) String() string {
	return "Card ID: " + c.card_id.String()
}

func (c *card) Dict() map[string]interface{} {
	return map[string]interface{}{
		"id":              c.id,
		"card_id":         c.card_id,
		"card_number":     c.card_number,
		"cvv_number":      c.cvv_number,
		"card_type":       c.card_type,
		"status":          c.status,
		"pin":             c.pin,
		"expiration_date": c.expiration_date,
		"salt_value":      c.salt_value,
		"created_date":    c.created_date,
		"updated_date":    c.updated_date,
	}
}

type Card struct{}

func (*Card) Create_card(card_type constants.CardType, pin string) (*card, error) {
	card_number, err := generate_card_number(card_type)
	if err != nil {
		return &card{}, fmt.Errorf("invalid card")
	}
	re := regexp2.MustCompile("^\\d{6}$", regexp2.None)
	if isMatch, _ := re.MatchString(pin); !isMatch {
		return &card{}, fmt.Errorf("invalid pin")
	}
	return &card{
		id:              uuid.New(),
		card_id:         uuid.New(),
		card_number:     card_number,
		cvv_number:      generate_cvv_number(),
		card_type:       card_type,
		status:          constants.Statuses.ACTIVE,
		pin:             pin,
		expiration_date: time.Now().AddDate(5, 0, 0),
		salt_value:      uuid.New(),
		created_date:    time.Now(),
		updated_date:    time.Now(),
	}, nil
}

func generate_card_number(card_type constants.CardType) (string, error) {
	value, err := card_type.String()
	if err != nil {
		return "", fmt.Errorf("no card number generated")
	}
	generic := ""
	for range [9]int{} {
		generic += strconv.Itoa(rand.Intn(9))

	}
	return value + generic, nil
}

func generate_cvv_number() string {
	generic := ""
	for range [6]int{} {
		generic += strconv.Itoa(rand.Intn(9))
	}
	return generic
}

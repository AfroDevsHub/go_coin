package models

import (
	"time"

	"github.com/dfunani/go_coin/lib/constants"
	"github.com/google/uuid"
)

type account struct {
	id           uuid.UUID
	account_id   uuid.UUID
	user_id      uuid.UUID
	status       constants.Status
	created_date time.Time
	updated_date time.Time
	// user_profiles = relationship(
	//     UserProfile, backref="Account", cascade="all, delete-orphan"
	// )
	// payment_profiles = relationship(
	//     PaymentProfile, backref="Account", cascade="all, delete-orphan"
	// )
	// settings_profile = relationship(
	//     SettingsProfile, backref="Account", cascade="all, delete-orphan"
	// )
}

func (a *account) String() string {
	return "Account ID: " + a.account_id.String()
}

func (a *account) Dict() map[string]interface{} {
	return map[string]interface{}{
		"id":           a.id,
		"account_id":   a.account_id,
		"user_id":      a.user_id,
		"status":       a.status,
		"created_date": a.created_date,
		"updated_date": a.updated_date,
	}
}

type Account struct{}

func (*Account) Create_account() *account {
	return &account{
		id:           uuid.New(),
		account_id:   uuid.New(),
		user_id:      uuid.New(),
		status:       constants.Statuses.NEW,
		created_date: time.Now(),
		updated_date: time.Now(),
	}
}

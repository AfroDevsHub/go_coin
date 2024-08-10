package users

import (
	"github.com/dfunani/go_coin/lib/constants"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID              uuid.UUID
	AccountID       uuid.UUID
	UserID          uuid.UUID
	Status          constants.Status
	PaymentProfiles PaymentProfile
	SettingsProfile SettingsProfile
	gorm.Model
}

func (*Account) TableName() string {
	return "users.accounts"
}

func (a *Account) String() string {
	return "Account ID: " + a.AccountID.String()
}

func (a *Account) Dict() map[string]interface{} {
	return map[string]interface{}{
		"ID":              a.ID,
		"AccountID":       a.AccountID,
		"UserID":          a.UserID,
		"Status":          a.Status,
		"PaymentProfiles": a.PaymentProfiles,
		"SettingsProfile": a.SettingsProfile.Dict(),
	}
}

type AccountSerialiser struct{}

func (*AccountSerialiser) Create_account() (*Account, error) {
	// settings_profiles := SettingsProfile{}
	// settings_profile, _ := settings_profiles.Create_settingsprofile(constants.UNVERIFIED, constants.UNVERIFIED, true)
	// user_profiles := UserProfile{}
	// user_profile, _ := user_profiles.Create_user_profile()

	return &Account{
		ID:              uuid.New(),
		AccountID:       uuid.New(),
		UserID:          uuid.New(),
		Status:          constants.Statuses.NEW,
		PaymentProfiles: PaymentProfile{},
		SettingsProfile: SettingsProfile{},
	}, nil
}

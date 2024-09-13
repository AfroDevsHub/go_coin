package serialisers

import (
	"encoding/json"
	"fmt"
	"time"

	DATABASE "github.com/dfunani/go_coin/database"
	CONSTANTS "github.com/dfunani/go_coin/lib/constants"
	USERS "github.com/dfunani/go_coin/models/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SettingsProfileSerialiser struct{}

func (*SettingsProfileSerialiser) ReadSettingsProfile(db *gorm.DB, settings_id uuid.UUID) (*USERS.SettingsProfile, error) {
	var settings []USERS.SettingsProfile
	db.Model(&USERS.SettingsProfile{}).Find(&settings, &USERS.SettingsProfile{SettingID: settings_id})
	if len(settings) != 1 {
		return &USERS.SettingsProfile{}, fmt.Errorf("invalid settings id")
	}
	return &settings[0], nil
}

func (*SettingsProfileSerialiser) CreateSettingsProfile(db *gorm.DB, account_id uuid.UUID, email_status CONSTANTS.Verification, communication_status CONSTANTS.Verification, mfa_enabled bool, profile_visibility_preference CONSTANTS.ProfileVisibility, data_sharing_preferences []CONSTANTS.DataSharingPreference, communication_preference CONSTANTS.Communication, location_tracking_enabled bool, cookies_enabled bool, theme_preference CONSTANTS.Theme) (string, error) {
	var mfa_last_used_date *time.Time = nil
	if mfa_enabled {
		var temp = time.Now()
		mfa_last_used_date = &temp
	}
	var valid_email_status string
	var valid_communication_status string
	var valid_profile_visibility_preference string
	var valid_data_sharing_preferences []byte
	var _data_sharing_preferences []string
	var valid_communication_preference string
	var valid_theme_preference string
	var err error
	if valid_email_status, err = email_status.String(); err != nil {
		return "", err
	}
	if valid_communication_status, err = communication_status.String(); err != nil {
		return "", err
	}
	if valid_profile_visibility_preference, err = profile_visibility_preference.String(); err != nil {
		return "", err
	}
	for _, data_sharing_preference := range data_sharing_preferences {
		var value string
		if value, err = data_sharing_preference.String(); err != nil {
			return "", err
		}
		_data_sharing_preferences = append(_data_sharing_preferences, value)
	}
	valid_data_sharing_preferences, _ = json.Marshal(_data_sharing_preferences)
	if valid_communication_preference, err = communication_preference.String(); err != nil {
		return "", err
	}
	if valid_theme_preference, err = theme_preference.String(); err != nil {
		return "", err
	}

	return DATABASE.CreateModel(db, &USERS.SettingsProfile{
		ID:                          uuid.New(),
		SettingID:                   uuid.New(),
		AccountID:                   account_id,
		EmailStatus:                 valid_email_status,
		CommunicationStatus:         valid_communication_status,
		MfaEnabled:                  mfa_enabled,
		MfaLastUsedDate:             mfa_last_used_date,
		ProfileVisibilityPreference: valid_profile_visibility_preference,
		DataSharingPreferences:      valid_data_sharing_preferences,
		CommunicationPreference:     valid_communication_preference,
		LocationTrackingEnabled:     location_tracking_enabled,
		CookiesEnabled:              cookies_enabled,
		ThemePreference:             valid_theme_preference,
	})
}

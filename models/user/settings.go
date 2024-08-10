package users

import (
	"time"

	"github.com/dfunani/go_coin/lib/constants"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SettingsProfile struct {
	gorm.Model
	ID                          uuid.UUID
	SettingsID                  uuid.UUID
	AccountID                   uuid.UUID
	EmailStatus                 constants.Verification
	CommunicationStatus         constants.Verification
	MfaEnabled                  bool
	MfaLastUsedDate             *time.Time
	ProfileVisibilityPreference constants.ProfileVisibility
	DataSharingPreferences      constants.DataSharingPreference
	CommunicationPreference     constants.Communication
	LocationTrackingEnabled     bool
	CookiesEnabled              bool
	ThemePreference             constants.Theme
}

func (l *SettingsProfile) TableName() string {
	return "users.settings_profiles"
}

func (s *SettingsProfile) String() string {
	return "Settings Profile ID: " + s.SettingsID.String()
}

func (s *SettingsProfile) Dict() map[string]interface{} {
	return map[string]interface{}{
		"id":                            s.ID,
		"settings_id":                   s.SettingsID,
		"account_id":                    s.AccountID,
		"email_status":                  s.EmailStatus,
		"communication_status":          s.CommunicationStatus,
		"mfa_enabled":                   s.MfaEnabled,
		"mfa_last_used_date":            s.MfaLastUsedDate,
		"profile_visibility_preference": s.ProfileVisibilityPreference,
		"data_sharing_preferences":      s.DataSharingPreferences,
		"communication_preference":      s.CommunicationPreference,
		"location_tracking_enabled":     s.LocationTrackingEnabled,
		"cookies_enabled":               s.CookiesEnabled,
		"theme_preference":              s.ThemePreference,
	}
}

type SettingsProfileSerialiser struct{}

func (*SettingsProfileSerialiser) Create_SettingsProfile(account_id uuid.UUID, email_status constants.Verification,
	communication_status constants.Verification, mfa_enabled bool, profile_visibility_preference constants.ProfileVisibility,
	data_sharing_preferences constants.DataSharingPreference,
	communication_preference constants.Communication,
	location_tracking_enabled bool,
	cookies_enabled bool,
	theme_preference constants.Theme) (*SettingsProfile, error) {
	var mfa_last_used_date *time.Time = nil
	if mfa_enabled {
		var temp = time.Now()
		mfa_last_used_date = &temp
	}
	return &SettingsProfile{
		ID:                          uuid.New(),
		SettingsID:                  uuid.New(),
		AccountID:                   account_id,
		EmailStatus:                 email_status,
		CommunicationStatus:         communication_status,
		MfaEnabled:                  mfa_enabled,
		MfaLastUsedDate:             mfa_last_used_date,
		ProfileVisibilityPreference: profile_visibility_preference,
		DataSharingPreferences:      data_sharing_preferences,
		CommunicationPreference:     communication_preference,
		LocationTrackingEnabled:     location_tracking_enabled,
		CookiesEnabled:              cookies_enabled,
		ThemePreference:             theme_preference,
	}, nil
}

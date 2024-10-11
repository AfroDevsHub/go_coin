package users

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SettingsProfile struct {
	ID                          uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	SettingID                   uuid.UUID  `json:"setting_id" gorm:"type:uuid;uniqueIndex;not null"`
	AccountID                   uuid.UUID  `json:"account_id" gorm:"type:uuid;uniqueIndex;not null"`
	EmailStatus                 string     `json:"email_status" gorm:"type:string"`
	CommunicationStatus         string     `json:"communication_status" gorm:"type:string"`
	MfaEnabled                  bool       `json:"mfa_enabled" gorm:"type:bool"`
	MfaLastUsedDate             *time.Time `json:"mfa_last_updated" gorm:"NOW()"`
	ProfileVisibilityPreference string     `json:"profile_visibility_preference" gorm:"type:string"`
	DataSharingPreferences      []byte     `json:"data_sharing_preferences"`
	CommunicationPreference     string     `json:"communication_preference" gorm:"type:string"`
	LocationTrackingEnabled     bool       `json:"location_tracking_enabled" gorm:"type:bool"`
	CookiesEnabled              bool       `json:"cookies_enabled" gorm:"type:bool"`
	ThemePreference             string     `json:"theme_preference" gorm:"type:string"`
	gorm.Model
}

func (*SettingsProfile) TableName() string {
	return "users.settings_profiles"
}

func (s *SettingsProfile) String() string {
	return "Settings Profile ID: " + s.SettingID.String()
}

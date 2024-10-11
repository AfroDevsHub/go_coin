package services

import (
	"time"

	CONSTANTS "github.com/dfunani/go_coin/lib/constants"
	"github.com/google/uuid"
)

type UserServiceRequest struct {
	UserID                      uuid.UUID
	Email                       string
	Password                    string
	Role                        CONSTANTS.Role
	AccountID                   uuid.UUID
	EmailStatus                 CONSTANTS.Verification
	CommunicationStatus         CONSTANTS.Verification
	MFAEnabled                  bool
	ProfileVisibilityPreference CONSTANTS.ProfileVisibility
	DataSharingPreferences      []CONSTANTS.DataSharingPreference
	CommunicationPreference     CONSTANTS.Communication
	LocationTrackingEnabled     bool
	CookiesEnabled              bool
	ThemePreference             CONSTANTS.Theme
	FirstName                   string
	LastName                    string
	Username                    string
	DateOfBirth                 time.Time
	Gender                      CONSTANTS.Gender
	ProfilePicture              string
	MobileNumber                string
	Country                     CONSTANTS.Country
	Language                    CONSTANTS.Language
	Biography                   string
	Occupation                  CONSTANTS.Occupation
	Interests                   CONSTANTS.Interest
	SocialMediaLinks            map[CONSTANTS.SocialMediaLink]string
	Status                      CONSTANTS.Status
	Location                    CONSTANTS.Country
	Device                      string
	Method                      CONSTANTS.LoginMethod
	Token                       string
}

type UserPaymentRequest struct {
	AccountID uuid.UUID
	CardType  CONSTANTS.CardType
	Pin       string
}

type UserProfileResponse struct {
	UserProfile     string
	SettingsProfile string
}

type UserResponse struct {
	UserID      string
	AccountID   string
	SettingsID  string
	UserProfile string
}

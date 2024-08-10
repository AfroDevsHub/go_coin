package users

import (
	"time"

	"github.com/dfunani/go_coin/lib/constants"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model
	ID               uuid.UUID
	ProfileID        uuid.UUID
	AccountID        uuid.UUID
	FirstName        string
	LastName         string
	Username         string
	DateOfBirth      time.Time
	Gender           constants.Gender
	ProfilePicture   string
	MobileNumber     string
	Country          constants.Country
	Language         constants.Language
	Biography        string
	Occupation       constants.Occupation
	Interests        constants.Interest
	SocialMediaLinks constants.SocialMediaLink
	Status           constants.Status
}

func (*UserProfile) TableName() string {
	return "users.user_profiles"
}

func (up *UserProfile) String() string {
	return "User Profile ID: " + up.ProfileID.String()
}

func (up *UserProfile) Dict() map[string]interface{} {
	return map[string]interface{}{
		"id":                 up.ID,
		"profile_id":         up.ProfileID,
		"account_id":         up.AccountID,
		"first_name":         up.FirstName,
		"last_name":          up.LastName,
		"username":           up.Username,
		"date_of_birth":      up.DateOfBirth,
		"gender":             up.Gender,
		"profile_picture":    up.ProfilePicture,
		"mobile_number":      up.MobileNumber,
		"country":            up.Country,
		"language":           up.Language,
		"biography":          up.Biography,
		"occupation":         up.Occupation,
		"interests":          up.Interests,
		"social_media_links": up.SocialMediaLinks,
		"status":             up.Status,
	}
}

type UserProfileSerialiser struct{}

func (*UserProfileSerialiser) Create_user_profile(first_name string, last_name string, username string, date_of_birth time.Time, gender constants.Gender, profile_picture string,
	mobile_number string,
	country constants.Country,
	language constants.Language,
	biography string,
	occupation constants.Occupation,
	interests constants.Interest,
	social_media_links constants.SocialMediaLink,
	status constants.Status) (*UserProfile, error) {
	return &UserProfile{
		ID:               uuid.New(),
		ProfileID:        uuid.New(),
		AccountID:        uuid.New(),
		FirstName:        first_name,
		LastName:         last_name,
		Username:         username,
		DateOfBirth:      date_of_birth,
		Gender:           gender,
		ProfilePicture:   profile_picture,
		MobileNumber:     mobile_number,
		Country:          country,
		Language:         language,
		Biography:        biography,
		Occupation:       occupation,
		Interests:        interests,
		SocialMediaLinks: social_media_links,
		Status:           status,
	}, nil
}

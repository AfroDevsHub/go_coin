package users

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserProfile struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	ProfileID        uuid.UUID `json:"profile_id" gorm:"type:uuid;not null;uniqueIndex"`
	AccountID        uuid.UUID `json:"account_id" gorm:"type:uuid;not null;uniqueIndex"`
	FirstName        string    `json:"first_name" gorm:"type:string"`
	LastName         string    `json:"last_name" gorm:"type:string"`
	Username         string    `json:"username" gorm:"type:string"`
	DateOfBirth      time.Time `json:"date_of_birth" gorm:"type:timestamp"`
	Gender           string    `json:"gender" gorm:"type:string"`
	ProfilePicture   string    `json:"profile_picture" gorm:"type:string"`
	MobileNumber     string    `json:"mobile_number" gorm:"type:string"`
	Country          string    `json:"country" gorm:"type:string"`
	Language         string    `json:"language" gorm:"type:string"`
	Biography        string    `json:"biography" gorm:"type:string"`
	Occupation       string    `json:"occupation" gorm:"type:string"`
	Interests        string    `json:"interests" gorm:"type:string"`
	SocialMediaLinks []byte    `json:"social_media_links"`
	Status           string    `json:"status" gorm:"type:string"`
	gorm.Model
}

func (*UserProfile) TableName() string {
	return "users.user_profiles"
}

func (up *UserProfile) String() string {
	return "User Profile ID: " + up.ProfileID.String()
}

package serialisers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dfunani/go_coin/database"
	"github.com/dfunani/go_coin/lib/constants"
	users "github.com/dfunani/go_coin/models/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserProfileSerialiser struct{}

func (*UserProfileSerialiser) ReadUserProfile(db *gorm.DB, profile_id uuid.UUID) (*users.UserProfile, error) {
	var user_profile []users.UserProfile
	db.Model(&users.UserProfile{}).Find(&user_profile, users.UserProfile{ProfileID: profile_id})
	if len(user_profile) != 1 {
		return &users.UserProfile{}, fmt.Errorf("invalid user profile id")
	}
	return &user_profile[0], nil
}

func (*UserProfileSerialiser) CreateUserProfile(db *gorm.DB, account_id uuid.UUID, first_name string, last_name string, username string, date_of_birth time.Time, gender constants.Gender, profile_picture string,
	mobile_number string,
	country constants.Country,
	language constants.Language,
	biography string,
	occupation constants.Occupation,
	interests constants.Interest,
	social_media_links map[constants.SocialMediaLink]string,
	status constants.Status) (string, error) {
	var valid_gender constants.TupleGender
	var valid_country constants.TupleCountry
	var valid_language constants.TupleLanguage
	var valid_occupation string
	var valid_interests string
	var valid_status string
	var valid_social_media_links = make(map[string]string)
	var err error

	if valid_gender, err = gender.String(); err != nil {
		return "", err
	}
	if valid_country, err = country.String(); err != nil {
		return "", err
	}
	if valid_language, err = language.String(); err != nil {
		return "", err
	}
	if valid_occupation, err = occupation.String(); err != nil {
		return "", err
	}
	if valid_interests, err = interests.String(); err != nil {
		return "", err
	}
	if valid_status, err = status.String(); err != nil {
		return "", err
	}
	for k, v := range social_media_links {
		regex, err := k.String()
		if err != nil {
			return "", fmt.Errorf("invalid social media links")
		}
		if ok, _ := regex.Regex.MatchString(v); !ok {
			return "", fmt.Errorf("invalid social media link")
		}
		valid_social_media_links[regex.Name] = v
	}
	s, _ := json.Marshal(valid_social_media_links)
	return database.CreateModel(db, &users.UserProfile{
		ID:               uuid.New(),
		ProfileID:        uuid.New(),
		AccountID:        account_id,
		FirstName:        first_name,
		LastName:         last_name,
		Username:         username,
		DateOfBirth:      date_of_birth,
		Gender:           valid_gender.Gender,
		ProfilePicture:   profile_picture,
		MobileNumber:     mobile_number,
		Country:          valid_country.Country,
		Language:         valid_language.Language,
		Biography:        biography,
		Occupation:       valid_occupation,
		Interests:        valid_interests,
		SocialMediaLinks: s,
		Status:           valid_status,
	})
}

func (*UserProfileSerialiser) UpdateUserProfile(db *gorm.DB, id uuid.UUID, user_profile_data map[string]interface{}) (string, error) {
	var user_profile users.UserProfile
	err := db.Model(&users.UserProfile{}).First(&user_profile, id).Error
	if err != nil {
		return "", err
	}
	db.Model(user_profile).Updates(user_profile_data)
	return user_profile.String(), nil
}

func (*UserProfileSerialiser) DeleteUserProfile(db *gorm.DB, id uuid.UUID) (string, error) {
	err := db.Delete(&users.UserProfile{}, id).Error
	if err != nil {
		return "", err
	}
	return id.String() + " Deleted", nil
}

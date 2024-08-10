package users

import (
	"fmt"

	"github.com/dfunani/go_coin/lib/constants"
	"github.com/dfunani/go_coin/models/warehouse"
	"github.com/dlclark/regexp2"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	UserID       uuid.UUID
	Email        string
	Password     string
	Status       string
	SaltValue    uuid.UUID
	Role         string
	LoginHistory []*warehouse.Loginhistory
	gorm.Model
}

func (*User) TableName() string {
	return "users.users"
}

func (u *User) String() string {
	return "User ID: " + u.UserID.String()
}

func (u *User) Dict() map[string]interface{} {
	return map[string]interface{}{
		"id":            u.ID,
		"userID":        u.UserID,
		"email":         u.Email,
		"password:":     u.Password,
		"status":        u.Status,
		"salt_value":    u.SaltValue,
		"role":          u.Role,
		"login_history": u.LoginHistory,
	}
}

type UserSerialiser struct{}

func (*UserSerialiser) Create_user(email string, password string) (*User, error) {
	user_id := uuid.New()
	err := validate_email(email)
	if err != nil {
		return &User{}, err
	}

	err = validate_password(password)
	if err != nil {
		return &User{}, err
	}

	status, _ := constants.Statuses.NEW.String()
	role, _ := constants.Roles.DEVELOPER.String()
	return &User{
		ID:           user_id,
		UserID:       uuid.New(),
		Email:        email,
		Password:     password,
		Status:       status,
		SaltValue:    uuid.New(),
		Role:         role,
		LoginHistory: nil,
	}, nil
}

func validate_email(email string) error {
	re := regexp2.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-_]+\.[a-zA-Z]{1,3}\.?[a-zA-Z]{2,3}$`, regexp2.IgnoreCase)
	if isMatch, _ := re.MatchString(email); !isMatch {
		return fmt.Errorf("invalid email")
	}
	return nil
}

func validate_password(password string) error {
	re := regexp2.MustCompile(`^(?=.*[a-zA-Z])(?=.*\d)(?=.*[!@#$%^&*()-_+=]).{8,}$`, regexp2.None)
	if isMatch, _ := re.MatchString(password); !isMatch {
		return fmt.Errorf("invalid password")
	}
	return nil
}

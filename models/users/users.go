package models

import (
	"fmt"
	"time"

	"github.com/dfunani/go_coin/lib/constants"
	"github.com/dfunani/go_coin/models/warehouse"
	"github.com/dlclark/regexp2"

	"github.com/google/uuid"
)

type user struct {
	id            uuid.UUID
	user_id       uuid.UUID
	email         string
	password      string
	created_date  time.Time
	updated_date  time.Time
	status        constants.Status
	salt_value    uuid.UUID
	role          constants.Role
	login_history map[string]interface{}
}

func (u *user) String() string {
	u.Dict()
	return "User ID: " + u.user_id.String()
}

func (l *user) Dict() map[string]interface{} {
	return map[string]interface{}{
		"id":            l.id,
		"user_id":       l.user_id,
		"email":         l.email,
		"password:":     l.password,
		"created_date:": l.created_date,
		"updated_date":  l.updated_date,
		"status":        l.status,
		"salt_value":    l.salt_value,
		"role":          l.role,
		"login_history": l.login_history,
	}
}

type User struct{}

func (*User) Create_user(email string, password string) (*user, error) {
	login := warehouse.LoginHistory{}
	re := regexp2.MustCompile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$", regexp2.IgnoreCase)
	if isMatch, _ := re.MatchString(email); !isMatch {
		return &user{}, fmt.Errorf("invalid email")
	}
	re = regexp2.MustCompile("^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[!@#$%^&*()_+\\-=\\[\\]{};':\"\\|,.<>\\/?]).{8,64}$", regexp2.IgnoreCase)
	if isMatch, _ := re.MatchString(password); !isMatch {
		return &user{}, fmt.Errorf("invalid password")
	}
	return &user{
		id:            uuid.New(),
		user_id:       uuid.New(),
		email:         email,
		password:      password,
		created_date:  time.Now(),
		updated_date:  time.Now(),
		status:        constants.Statuses.NEW,
		salt_value:    uuid.New(),
		role:          constants.Roles.DEVELOPER,
		login_history: login.Create_login_history().Dict(),
	}, nil
}

package warehouse

import (
	"time"

	"github.com/dfunani/go_coin/lib/constants"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Loginhistory struct {
	gorm.Model
	ID                  uuid.UUID `gorm:"primaryKey"`
	LoginID             uuid.UUID
	SessionID           uuid.UUID
	UserID              uuid.UUID
	LoginDate           time.Time
	LoginLocation       constants.Country
	LoginDevice         string
	LoginMethod         constants.LoginMethod
	LoggedIn            bool
	LogoutDate          *time.Time
	AuthenticationToken string
}

func (l *Loginhistory) TableName() string {
	return "warehouse.login_histories"
}

func (l *Loginhistory) String() string {
	return "Login History ID: " + l.LoginID.String()
}

func (l *Loginhistory) Dict() map[string]interface{} {
	return map[string]interface{}{
		"ID":                  l.ID,
		"LoginID":             l.LoginID,
		"SessionID":           l.SessionID,
		"UserID":              l.UserID,
		"LoginDate":           l.LoginDate,
		"LoginLocation":       l.LoginLocation,
		"LoginDevice":         l.LoginDevice,
		"LoginMethod":         l.LoginMethod,
		"LoggedIn":            l.LoggedIn,
		"LogoutDate":          l.LogoutDate,
		"AuthenticationToken": l.AuthenticationToken,
	}
}

type LoginHistorySerialiser struct{}

func (*LoginHistorySerialiser) Create_login_history(user_id uuid.UUID, location constants.Country, device string, method constants.LoginMethod, token string) *Loginhistory {
	return &Loginhistory{
		ID:                  uuid.New(),
		LoginID:             uuid.New(),
		SessionID:           uuid.New(),
		UserID:              user_id,
		LoginDate:           time.Now(),
		LoginLocation:       location,
		LoginDevice:         device,
		LoginMethod:         method,
		LoggedIn:            true,
		LogoutDate:          nil,
		AuthenticationToken: token,
	}
}

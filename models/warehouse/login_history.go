package warehouse

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Loginhistory struct {
	ID                  uuid.UUID  `gorm:"primaryKey"`
	LoginID             uuid.UUID  `json:"login_id" gorm:"type:uuid;uniqueIndex;not null"`
	SessionID           uuid.UUID  `json:"session_id" gorm:"type:uuid;uniqueIndex;not null"`
	UserID              uuid.UUID  `json:"user_id" gorm:"type:uuid;uniqueIndex;not null"`
	LoginDate           time.Time  `json:"login_date" gorm:"type:timestamp;not null"`
	LoginLocation       string     `json:"login_location" gorm:"type:string;not null"`
	LoginDevice         string     `json:"login_device" gorm:"type:string;not null"`
	LoginMethod         string     `json:"login_method" gorm:"type:string;not null"`
	LoggedIn            bool       `json:"logged_in" gorm:"type:boolean;not null"`
	LogoutDate          *time.Time `json:"logout_date" gorm:"type:timestamp"`
	AuthenticationToken string     `json:"authentication_token" gorm:"type:string;not null"`
	gorm.Model
}

func (l *Loginhistory) TableName() string {
	return "warehouse.login_histories"
}

func (l *Loginhistory) String() string {
	return "Login History ID: " + l.LoginID.String()
}

/*Models: Defines User Related Model Schemas.*/
package users

import (
	WAREHOUSE "github.com/dfunani/go_coin/models/warehouse"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID                `json:"id" gorm:"type:uuid;primary_key"`
	UserID         uuid.UUID                `json:"user_id" gorm:"type:uuid;uniqueIndex;not null"`
	Email          string                   `json:"email" gorm:"type:string;uniqueIndex;not null"`
	Password       string                   `json:"password" gorm:"type:string;not null"`
	Status         string                   `json:"status" gorm:"type:string;not null"`
	SaltValue      uuid.UUID                `json:"salt_value" gorm:"type:uuid;not null"`
	Role           string                   `json:"role" gorm:"type:string;not null"`
	Account        Account                  `gorm:"foreignKey:UserID;references:ID"`
	LoginHistories []WAREHOUSE.Loginhistory `gorm:"foreignKey:UserID;references:ID"`
	gorm.Model
}

func (*User) TableName() string {
	return "users.users"
}

func (u *User) String() string {
	return "User ID: " + u.UserID.String()
}

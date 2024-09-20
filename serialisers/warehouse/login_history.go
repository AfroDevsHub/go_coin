package serialisers

import (
	"fmt"
	"time"

	"github.com/dfunani/go_coin/database"
	"github.com/dfunani/go_coin/lib/constants"
	"github.com/dfunani/go_coin/models/warehouse"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoginHistorySerialiser struct{}

func (*LoginHistorySerialiser) ReadLoginHistory(db *gorm.DB, card_id uuid.UUID) (*warehouse.Loginhistory, error) {
	var card []warehouse.Loginhistory
	db.Model(&warehouse.Loginhistory{}).Find(&card, &warehouse.Loginhistory{LoginID: card_id})
	if len(card) != 1 {
		return &warehouse.Loginhistory{}, fmt.Errorf("invalid card id")
	}
	return &card[0], nil
}

func (*LoginHistorySerialiser) Create_login_history(db *gorm.DB, user_id uuid.UUID, location constants.Country, device string, method constants.LoginMethod, token string) (string, error) {
	var valid_location constants.TupleCountry
	var valid_method string
	var err error
	if valid_location, err = location.String(); err != nil {
		return "", err
	}
	if valid_method, err = method.String(); err != nil {
		return "", err
	}
	return database.CreateModel(db, &warehouse.Loginhistory{
		ID:                  uuid.New(),
		LoginID:             uuid.New(),
		SessionID:           uuid.New(),
		UserID:              user_id,
		LoginDate:           time.Now(),
		LoginLocation:       valid_location.Country,
		LoginDevice:         device,
		LoginMethod:         valid_method,
		LoggedIn:            true,
		LogoutDate:          nil,
		AuthenticationToken: token,
	})
}

func (*LoginHistorySerialiser) UpdateLoginHistory(db *gorm.DB, id uuid.UUID, login_history_data map[string]interface{}) (string, error) {
	var login_history warehouse.Loginhistory
	err := db.Model(&warehouse.Loginhistory{}).First(&login_history, id).Error
	if err != nil {
		return "", err
	}
	db.Model(login_history).Updates(login_history_data)
	return login_history.String(), nil
}

func (*LoginHistorySerialiser) DeleteLoginHistory(db *gorm.DB, id uuid.UUID) (string, error) {
	err := db.Delete(&warehouse.Loginhistory{}, id).Error
	if err != nil {
		return "", err
	}
	return id.String() + " Deleted", nil
}

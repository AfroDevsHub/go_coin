package serialisers

import (
	"fmt"

	DATABASE "github.com/dfunani/go_coin/database"
	CONSTANTS "github.com/dfunani/go_coin/lib/constants"
	USERS "github.com/dfunani/go_coin/models/user"
	users "github.com/dfunani/go_coin/models/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountSerialiser struct{}

/*
Description:

	Creates an Account Object and Persists the Account in the Database.

Parameters:

 1. db - Database Connection.

 2. user_id - Unique Identifier for a User. # Example: User ID: 1234-5678-9ABC-DEFG

Returns:

	String Representation of a Valid Account Model. # Example: Account ID: 1234-5678-9ABC-DEFG
	String Representation of an Error Object. # Example: Error String.
*/

func (*AccountSerialiser) ReadAccount(db *gorm.DB, account_id uuid.UUID) (*USERS.Account, error) {
	var account []USERS.Account
	db.Model(&USERS.Account{}).Find(&account, USERS.Account{AccountID: account_id})
	if len(account) != 1 {
		return &USERS.Account{}, fmt.Errorf("invalid account id")
	}
	return &account[0], nil
}

func (*AccountSerialiser) CreateAccount(db *gorm.DB, user_id uuid.UUID) (string, error) {
	status, err := CONSTANTS.Statuses.NEW.String()
	if err != nil {
		return "", err
	}
	return DATABASE.CreateModel(db, &USERS.Account{
		ID:        uuid.New(),
		AccountID: uuid.New(),
		UserID:    user_id,
		Status:    status,
	})
}

func (*AccountSerialiser) UpdateAccount(db *gorm.DB, id uuid.UUID, account_data map[string]interface{}) (string, error) {
	var account USERS.Account
	err := db.Model(&USERS.Account{}).First(&account, id).Error
	if err != nil {
		return "", err
	}
	db.Model(account).Updates(account_data)
	return account.String(), nil
}

func (*AccountSerialiser) DeleteAccount(db *gorm.DB, id uuid.UUID) (string, error) {
	err := db.Delete(&users.User{}, id).Error
	if err != nil {
		return "", err
	}
	return id.String() + " Deleted", nil
}

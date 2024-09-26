/*Serialisers: Manages User Related Models, by performing CRUD operations.*/
package serialisers

import (
	"fmt"
	"log"
	"os"

	DATABASE "github.com/dfunani/go_coin/database"
	CONSTANTS "github.com/dfunani/go_coin/lib/constants"
	"github.com/dfunani/go_coin/lib/utils"
	UTILS "github.com/dfunani/go_coin/lib/utils"
	USERS "github.com/dfunani/go_coin/models/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserSerialiser struct{}

/*
Description:

	Creates a User Object and Persists the User in the Database.

Parameters:

 1. db - Database Connection.

 2. email - Valid User Email.

 3. password - Valid User Password.

 4. role - Valid User Role.

Returns:

	String Representation of a Valid User Model. # Example: User ID: 1234-5678-9ABC-DEFG
	String Representation of an Error Object. # Example: Error String.
*/

func (*UserSerialiser) ReadUser(db *gorm.DB, user_id uuid.UUID) (*USERS.User, error) {
	var users []USERS.User
	db.Model(&USERS.User{}).Find(&users, USERS.User{UserID: user_id})
	if len(users) != 1 {
		return &USERS.User{}, fmt.Errorf("invalid user id")
	}
	return &users[0], nil
}

func (*UserSerialiser) CreateUser(db *gorm.DB, email string, password string, role CONSTANTS.Role) (string, error) {
	err := UTILS.ValidateEmail(email)
	if err != nil {
		return "", err
	}
	err = UTILS.ValidatePassword(password)
	if err != nil {
		return "", err
	}
	if role == CONSTANTS.Roles.ADMIN {
		return "", fmt.Errorf("invalid role")
	}

	var valid_status string
	if valid_status, err = CONSTANTS.Statuses.NEW.String(); err != nil {
		return "", err
	}
	var valid_role string
	if valid_role, err = role.String(); err != nil {
		return "", err
	}
	secret, _ := os.LookupEnv("SECRET")
	encrypted_email, err := utils.Encrypt([]byte(email), []byte(utils.GenerateHash(secret)))
	if err != nil {
		return "", err
	}
	hash_password := utils.GenerateHash(password)
	log.Println(encrypted_email, hash_password)

	return DATABASE.CreateModel(db, &USERS.User{
		ID:        uuid.New(),
		UserID:    uuid.New(),
		Email:     encrypted_email,
		Password:  hash_password,
		Status:    valid_status,
		SaltValue: uuid.New(),
		Role:      valid_role,
	})
}

func (*UserSerialiser) UpdateUser(db *gorm.DB, id uuid.UUID, user_data map[string]interface{}) (string, error) {
	var user USERS.User
	err := db.Model(&USERS.User{}).First(&user, id).Error
	if err != nil {
		return "", err
	}
	db.Model(user).Updates(user_data)
	return user.String(), nil
}

func (*UserSerialiser) DeleteUser(db *gorm.DB, id uuid.UUID) (string, error) {
	err := db.Delete(&USERS.User{}, id).Error
	if err != nil {
		return "", err
	}
	return id.String() + " Deleted", nil
}

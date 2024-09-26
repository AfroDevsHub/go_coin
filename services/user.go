package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dfunani/go_coin/lib/utils"
	UTILS "github.com/dfunani/go_coin/lib/utils"
	USERS "github.com/dfunani/go_coin/models/user"
	WAREHOUSE "github.com/dfunani/go_coin/models/warehouse"
	USER_SERIALISERS "github.com/dfunani/go_coin/serialisers/user"
	WAREHOUSE_SERIALISERS "github.com/dfunani/go_coin/serialisers/warehouse"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func registerUser(db *gorm.DB, request *UserServiceRequest) (*USERS.User, error) {
	user_serialiser := USER_SERIALISERS.UserSerialiser{}
	user, err := user_serialiser.CreateUser(db, request.Email, request.Password, request.Role)
	if err != nil {
		return &USERS.User{}, fmt.Errorf("error - user - %s", err)
	}
	user_id, err := UTILS.GetModelID(user)
	if err != nil {
		return &USERS.User{}, fmt.Errorf("error - user id - %s", err)
	}
	get_user, err := user_serialiser.ReadUser(db, uuid.MustParse(user_id))
	if err != nil {
		return &USERS.User{}, err
	}
	return get_user, nil
}

func registerAccount(db *gorm.DB, request *UserServiceRequest) (*USERS.Account, error) {
	account_serialiser := USER_SERIALISERS.AccountSerialiser{}
	account, err := account_serialiser.CreateAccount(db, request.UserID)
	if err != nil {
		return &USERS.Account{}, fmt.Errorf("error - account - %s", err)
	}
	account_id, err := UTILS.GetModelID(account)
	if err != nil {
		return &USERS.Account{}, fmt.Errorf("error - account id - %s", err)
	}
	get_account, err := account_serialiser.ReadAccount(db, uuid.MustParse(account_id))
	if err != nil {
		return &USERS.Account{}, err
	}
	return get_account, nil
}

func registerProfile(db *gorm.DB, request *UserServiceRequest) (*UserProfileResponse, error) {
	settings_serialiser := USER_SERIALISERS.SettingsProfileSerialiser{}
	settings, err := settings_serialiser.CreateSettingsProfile(db, request.AccountID, request.EmailStatus, request.CommunicationStatus, request.MFAEnabled, request.ProfileVisibilityPreference, request.DataSharingPreferences, request.CommunicationPreference, request.LocationTrackingEnabled, request.CookiesEnabled, request.ThemePreference)
	if err != nil {
		return &UserProfileResponse{}, fmt.Errorf("error - settings - %s", err)
	}
	profile := UserProfileResponse{}
	profile.SettingsProfile = settings

	user_profile_serialiser := USER_SERIALISERS.UserProfileSerialiser{}
	user_profile, err := user_profile_serialiser.CreateUserProfile(db, request.AccountID, request.FirstName, request.LastName, request.Username, request.DateOfBirth, request.Gender, request.ProfilePicture, request.MobileNumber, request.Country, request.Language, request.Biography, request.Occupation, request.Interests, request.SocialMediaLinks, request.Status)
	if err != nil {
		return &UserProfileResponse{}, fmt.Errorf("error - user profile - %s", err)
	}
	profile.UserProfile = user_profile
	return &profile, nil
}

func Register(db *gorm.DB, request *UserServiceRequest) (*UserResponse, error) {
	user, err := registerUser(db, request)
	if err != nil {
		return &UserResponse{}, err
	}
	request.UserID = user.ID
	log.Println("INFO - User ID -", user.ID)

	account, err := registerAccount(db, request)
	if err != nil {
		return &UserResponse{}, err
	}
	request.AccountID = account.ID
	log.Println("INFO - Account ID -", account.ID)

	profile, err := registerProfile(db, request)
	if err != nil {
		return &UserResponse{}, err
	}
	log.Println("INFO -", profile.SettingsProfile)
	log.Println("INFO -", profile.UserProfile)

	return &UserResponse{UserID: user.String(), AccountID: account.String(), SettingsID: profile.SettingsProfile, UserProfile: profile.UserProfile}, nil

}

func RegisterPayment(db *gorm.DB, request *UserPaymentRequest) (string, error) {
	payment_serialiser := USER_SERIALISERS.PaymentProfileSerialiser{}
	payment, err := payment_serialiser.CreatePaymentProfile(db, request.AccountID, request.CardType, request.Pin)
	if err != nil {
		return "", err
	}
	log.Println("INFO -", payment)
	return payment, nil
}

func Login(db *gorm.DB, request *UserServiceRequest) (string, error) {
	var user []USERS.User
	secret, _ := os.LookupEnv("SECRET")
	email, _ := utils.Encrypt([]byte(request.Email), []byte(utils.GenerateHash(secret)))
	password := utils.GenerateHash(request.Password)
	db.Model(&USERS.User{}).Find(&user, USERS.User{Email: email, Password: password})
	if len(user) != 1 {
		return "", fmt.Errorf("invalid user login")
	}
	login_history_serialiser := WAREHOUSE_SERIALISERS.LoginHistorySerialiser{}
	login, err := login_history_serialiser.Create_login_history(db, user[0].ID, request.Location, request.Device, request.Method, request.Token)
	if err != nil {
		return "", fmt.Errorf("error - login history - %s", err)
	}
	return login, nil
}

func Logout(db *gorm.DB, id uuid.UUID) (string, error) {
	var logins []WAREHOUSE.Loginhistory
	db.Model(&WAREHOUSE.Loginhistory{}).Find(&logins, WAREHOUSE.Loginhistory{ID: id, LoggedIn: true})
	for _, login := range logins {
		db.Model(&login).Updates(map[string]interface{}{"LoggedIn": false, "LoginDate": time.Now()})
	}
	return fmt.Sprintln("Logged Out Successfully:", id), nil
}

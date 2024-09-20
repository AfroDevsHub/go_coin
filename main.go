/*Main: Application Entry Point.*/

package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"

	DATABASE "github.com/dfunani/go_coin/database"
	users "github.com/dfunani/go_coin/models/user"
	"github.com/dfunani/go_coin/models/warehouse"
	serialisers "github.com/dfunani/go_coin/serialisers/user"
	w_serialisers "github.com/dfunani/go_coin/serialisers/warehouse"

	"github.com/dfunani/go_coin/lib/constants"
	"github.com/dfunani/go_coin/lib/utils"
	"github.com/dfunani/go_coin/services"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := DATABASE.SetupDatabase()
	number := rand.Intn(uuid.Max.ClockSequence())
	user_request := services.UserServiceRequest{
		Email:                  "delali" + strconv.Itoa(number) + "@funani.co.za",
		Password:               "DF@8_letters",
		DataSharingPreferences: []constants.DataSharingPreference{constants.ACCOUNT, constants.PROFILE, constants.SETTINGS},
	}

	user_response, _ := services.Register(db, &user_request)
	log.Println(user_response)
	login_response, _ := services.Login(db, &user_request)
	log.Println(login_response)

	var logout warehouse.Loginhistory
	id, _ := utils.GetModelID(login_response)
	db.Model(&warehouse.Loginhistory{}).First(&logout, warehouse.Loginhistory{LoginID: uuid.MustParse(id)})
	logout_response, _ := services.Logout(db, logout.ID)
	log.Println(logout_response)

	var random_user users.User
	var random_account users.Account
	db.Model(&users.User{}).First(&random_user)
	db.Model(&users.Account{}).First(&random_account, users.Account{UserID: random_user.ID})
	payment_request := services.UserPaymentRequest{AccountID: random_account.ID, Pin: "123456", CardType: constants.SAVINGS}
	payment_response, _ := services.RegisterPayment(db, &payment_request)
	log.Println(payment_response)

	user_serialiser := serialisers.UserSerialiser{}
	user_id, _ := utils.GetModelID(user_response.UserID)
	user, _ := user_serialiser.ReadUser(db, uuid.MustParse(user_id))
	log.Println(user)

	account_serialiser := serialisers.AccountSerialiser{}
	account_id, _ := utils.GetModelID(random_account.String())
	account, _ := account_serialiser.ReadAccount(db, uuid.MustParse(account_id))
	log.Println(account)

	settings_serialiser := serialisers.SettingsProfileSerialiser{}
	settings_id, _ := utils.GetModelID(user_response.SettingsID)
	settings, _ := settings_serialiser.ReadSettingsProfile(db, uuid.MustParse(settings_id))
	log.Println(settings)

	user_profile_serialiser := serialisers.UserProfileSerialiser{}
	profile_id, _ := utils.GetModelID(user_response.UserProfile)
	user_profile, _ := user_profile_serialiser.ReadUserProfile(db, uuid.MustParse(profile_id))
	log.Println(user_profile)

	payment_serialiser := serialisers.PaymentProfileSerialiser{}
	payment_id, _ := utils.GetModelID(payment_response)
	payment, _ := payment_serialiser.ReadPaymentProfile(db, uuid.MustParse(payment_id))
	log.Println(payment)

	login_history_serialiser := w_serialisers.LoginHistorySerialiser{}
	login_id, _ := utils.GetModelID(login_response)
	login_history, _ := login_history_serialiser.ReadLoginHistory(db, uuid.MustParse(login_id))
	log.Println(login_history)

	var random_card warehouse.Card
	card_serialiser := w_serialisers.CardSerialiser{}
	db.Model(&warehouse.Card{}).First(&random_card, warehouse.Card{ID: payment.CardID})
	card, _ := card_serialiser.ReadCard(db, random_card.CardID)
	log.Println(card)

	res, _ := user_serialiser.DeleteUser(db, uuid.MustParse(("97a20ff2-4b35-4b1e-8403-e9e9de019ad1")))
	log.Println((res))

	os.Exit(0)
}

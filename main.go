/*Main: Application Entry Point.*/

package main

import (
	"log"
	"os"
	"strconv"

	DATABASE "github.com/dfunani/go_coin/database"
	"github.com/dfunani/go_coin/lib/constants"
	serialisers "github.com/dfunani/go_coin/serialisers/user"
	"github.com/dfunani/go_coin/services"
	"github.com/google/uuid"
	"golang.org/x/exp/rand"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := DATABASE.SetupDatabase()
	number := rand.Intn(uuid.Max.ClockSequence())
	email := "delali" + strconv.Itoa(number) + "@funani.co.za"
	password := "DF@8_letters"
	log.Println(email, password)
	user_serialiser := serialisers.UserSerialiser{}
	user_serialiser.CreateUser(db, email, password, constants.DEVELOPER)
	request := services.UserServiceRequest{
		Email:    email,
		Password: password,
	}
	log.Println(services.Login(db, &request))

	os.Exit(0)
}

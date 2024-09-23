/*Main: Application Entry Point.*/

package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"

	DATABASE "github.com/dfunani/go_coin/database"
	"github.com/dfunani/go_coin/lib/constants"
	"github.com/dfunani/go_coin/lib/utils"
	serialisers "github.com/dfunani/go_coin/serialisers/user"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := DATABASE.SetupDatabase()
	number := rand.Intn(uuid.Max.ClockSequence())
	email := "delali" + strconv.Itoa(number) + "@funani.co.za"
	password := "DF@8_letters"

	user_serialiser := serialisers.UserSerialiser{}
	user_serialiser.CreateUser(db, email, password, constants.DEVELOPER)
	log.Println(utils.GenerateHash("Bye"))

	// _ := uuid.NewString()
	encrypted_data, _ := utils.Encrypt([]byte("Bye"), "salt")
	log.Println(encrypted_data)
	log.Println(utils.Decrypt([]byte(encrypted_data), "salt"))

	os.Exit(0)
}

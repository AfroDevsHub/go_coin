package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dfunani/go_coin/database"
	MODELS "github.com/dfunani/go_coin/models"
	USERS "github.com/dfunani/go_coin/models/user"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var models = []MODELS.Model{
	&USERS.User{},
	&USERS.Account{},
	// &WAREHOUSE.Loginhistory{},
	// &WAREHOUSE.Card{},
	// &USERS.UserProfile{},
	// &USERS.SettingsProfile{},
}

func Setup_database(auto_migrate bool) (*gorm.DB, error) {
	db, err := database.Connection()

	if auto_migrate {
		for _, model := range models {
			response := database.Migrate(db, model)
			if response != nil {
				fmt.Println("Could not Automigrate:", response)
			}
		}
	}
	return db, err
}

func Create_model(db *gorm.DB, model MODELS.Model) (string, error) {
	var err = db.Create(model).Error

	if err != nil {
		return "", err
	}
	return model.String(), nil
}

func main() {
	godotenv.Load()

	db, err := Setup_database(true)
	if err != nil {
		log.Fatalln("Error Setting Up Database Connection:", err)
	}

	user_serialiser := USERS.UserSerialiser{}
	users, err := user_serialiser.Create_user("delali@funani.co.za", "DF@8_letters")
	if err != nil {
		log.Fatalln("Error Creating User:", err)
	}

	response, _ := Create_model(db, users)
	fmt.Println(response, "was created.")
	os.Exit(0)
}

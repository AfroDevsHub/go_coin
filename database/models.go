package database

import (
	"os"

	MODELS "github.com/dfunani/go_coin/models"
	USERS "github.com/dfunani/go_coin/models/user"
	WAREHOUSE "github.com/dfunani/go_coin/models/warehouse"
)

type DB struct {
	HOST     string
	USER     string
	PASSWORD string
	NAME     string
	PORT     string
}

type migrations struct {
	Success []string
	Failed  []string
}

func (db *DB) configure() *DB {
	return &DB{
		HOST:     os.Getenv("DB_HOST"),
		USER:     os.Getenv("DB_USER"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
		NAME:     os.Getenv("DB_NAME"),
		PORT:     os.Getenv("DB_PORT"),
	}
}

var models = []MODELS.Model{
	&USERS.User{},
	&USERS.Account{},
	&USERS.PaymentProfile{},
	&WAREHOUSE.Loginhistory{},
	&WAREHOUSE.Card{},
	&USERS.UserProfile{},
	&USERS.SettingsProfile{},
}

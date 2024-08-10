package database

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	HOST     string
	USER     string
	PASSWORD string
	NAME     string
	PORT     string
}

func (db *DB) Configure() *DB {
	return &DB{
		HOST:     os.Getenv("DB_HOST"),
		USER:     os.Getenv("DB_USER"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
		NAME:     os.Getenv("DB_NAME"),
		PORT:     os.Getenv("DB_PORT"),
	}
}

func generate_schemas(connection *gorm.DB) error {
	env_schemas := os.Getenv("DB_SCHEMA")
	schemas := strings.Split(env_schemas, ",")

	for _, schema := range schemas {
		sql := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", schema)
		err := connection.Exec(sql).Error
		if err != nil {
			return err
		}
	}
	return nil

}

func Connection() (*gorm.DB, error) {
	db := DB{}
	dbConfig := db.Configure()
	var CONNECTION_STRING = "host=" + dbConfig.HOST + " user=" + dbConfig.USER + " password=" + dbConfig.PASSWORD + " dbname=" + dbConfig.NAME + " port=" + dbConfig.PORT
	var connection, err = gorm.Open(postgres.Open(CONNECTION_STRING), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("error connecting to database. connection closed. %s", err)
	}
	if generate_schemas(connection) != nil {
		return nil, fmt.Errorf("error creating schemas. connection closed")
	}

	return connection, err
}

func Migrate(db *gorm.DB, model interface{}) error {
	return db.AutoMigrate(&model)
}

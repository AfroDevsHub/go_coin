package database

import (
	"log"

	MODELS "github.com/dfunani/go_coin/models"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	log.Println("INFO - Connecting to Database.")
	db, err := connect()
	if err != nil {
		log.Fatalln("Error - Database Connection Setup Failed:", err)
	}

	response := migrate(db)
	built_string := buildConnectionResponse(response)

	log.Println(built_string)
	log.Println("INFO - Database Connection Complete.")
	return db
}

func CreateModel(db *gorm.DB, model MODELS.Model) (string, error) {
	var err = db.Create(model).Error

	if err != nil {
		return "Model Not Created.", err
	}
	return model.String(), nil
}

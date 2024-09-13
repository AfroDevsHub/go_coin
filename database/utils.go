package database

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func generateSchemas(connection *gorm.DB) error {
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

func connect() (*gorm.DB, error) {
	db := DB{}
	dbConfig := db.configure()
	var CONNECTION_STRING = "host=" + dbConfig.HOST + " user=" + dbConfig.USER + " password=" + dbConfig.PASSWORD + " dbname=" + dbConfig.NAME + " port=" + dbConfig.PORT
	var connection, err = gorm.Open(postgres.Open(CONNECTION_STRING), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("error connecting to database. connection closed. %s", err)
	}
	if generateSchemas(connection) != nil {
		return nil, fmt.Errorf("error creating schemas. connection closed")
	}

	return connection, err
}

func migrate(db *gorm.DB) migrations {
	successful_migrations := []string{}
	failed_migrations := []string{}
	for _, model := range models {
		err := db.AutoMigrate(&model)
		if err != nil {
			failed_migrations = append(failed_migrations, model.TableName())
		} else {
			successful_migrations = append(successful_migrations, model.TableName())
		}
	}
	return migrations{successful_migrations, failed_migrations}
}

func buildConnectionResponse(migrations migrations) string {
	var length_migrations = len(migrations.Success) + len(migrations.Failed)
	var success = strings.Join(migrations.Success, "\n    ")
	var failed = strings.Join(migrations.Failed, "\n    ")

	var header = fmt.Sprintf("INFO - Auto-Migrations - Total of %v Migrations", length_migrations)
	header = fmt.Sprintf("%s (%v - Successful and %v - Failed)\n", header, len(migrations.Success), len(migrations.Failed))
	var success_string = fmt.Sprintf("1. Successful Migrations\n    %s\n", success)
	var failed_string = fmt.Sprintf("2. Failed Migrations\n    %s", failed)

	return fmt.Sprintln(header, success_string, failed_string)
}

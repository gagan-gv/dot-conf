package database

import (
	"dot_conf/configs"
	"dot_conf/models"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Initialize() error {
	var err error

	config := configs.NewDatabaseConfig()
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DbName, config.SslMode,
	)
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return err
	}

	log.Info("Initialized the db successfully")

	err = db.AutoMigrate(
		&models.App{},
		&models.User{},
		&models.Config{},
		&models.Company{},
	)

	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}

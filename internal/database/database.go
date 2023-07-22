package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/byte3/bookclub/backend/config"
	"github.com/byte3/bookclub/backend/internal/models"
)

var Db *gorm.DB

func Initialize(config *config.Config) {
	var err error
	dsn := "host=localhost user=rxored dbname=bookclub port=9920 sslmode=disable TimeZone=Asia/Colombo"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	if config.Environment == "production" {
		Db.Logger.LogMode(logger.Error)
	}

	if err = Db.AutoMigrate(
		// migrating models
		&models.Author{},
		&models.Book{},
		&models.BookAvailability{},
		&models.BookGenre{},
		&models.BookRequest{},
		&models.BookRequestAccepted{},
		&models.BookRequestDeclined{},
		&models.BookRequestStatus{},
		&models.BookReturns{},
		&models.ISBNVersion{},
		&models.User{},
		&models.UserAddressInfo{},
		&models.UserContactInfo{},
	); err != nil {
		log.Fatal(err.Error())
	}
}

func SelectAllUsers() ([]models.User, error) {
	users := []models.User{}
	result := Db.Find(&users)
	return users, result.Error
}

func SelectUserbyName(username string) (models.User, error) {
	var user models.User
	result := Db.Where("username = ?", username).First(&user)
	return user, result.Error
}

func InsertUser(user *models.User) (int, error) {
	result := Db.Create(&user)

	return int(result.RowsAffected), result.Error
}

func GetUserDetails(id int) (models.User, error) {
	var user models.User
	result := Db.First(&user, id)
	return user, result.Error
}

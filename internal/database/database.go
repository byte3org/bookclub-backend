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
		&models.BookRequestStatus{}, &models.BookReturns{},
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
	user := models.User{}
	result := Db.Where("username = ?", username).First(&user)
	return user, result.Error
}

func InsertUser(user *models.User) (int, error) {
	result := Db.Create(user)

	return int(result.RowsAffected), result.Error
}

func SelectUserById(id int) (models.User, error) {
	user := models.User{}
	result := Db.First(&user, id)
	return user, result.Error
}

func SelectAllRequests() ([]models.BookRequest, error) {
	reqs := []models.BookRequest{}
	result := Db.Find(&reqs)
	return reqs, result.Error
}

func InsertRequest(req *models.BookRequest) (int, error) {
	result := Db.Create(req)
	return int(result.RowsAffected), result.Error
}

func SelectAllAcceptedRequests() ([]models.BookRequest, error) {
	reqs := []models.BookRequest{}
	results := Db.Where("request_status = accepted").Find(&reqs)

	return reqs, results.Error
}

func SelectAllPendingRequests() ([]models.BookRequest, error) {
	reqs := []models.BookRequest{}
	results := Db.Where("request_status = pending").Find(&reqs)

	return reqs, results.Error
}

func SelectRequestById(id int) (models.BookRequest, error) {
	req := models.BookRequest{}
	result := Db.First(&req, id)
	return req, result.Error
}

func DeleteRequestById(id int) error {
	var req models.BookRequest
	req.Id = id
	result := Db.Delete(&req)
	return result.Error
}

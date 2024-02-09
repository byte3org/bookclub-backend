package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/byte3/bookclub/backend/config"
	"github.com/byte3/bookclub/backend/internal/models"
	"github.com/google/uuid"
)

var Db *gorm.DB

func Initialize(config *config.Config) {
	var err error
	println("database is", config.DbString)
	Db, err = gorm.Open(postgres.Open(config.DbString), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	if config.Environment == "production" {
		Db.Logger.LogMode(logger.Error)
	}

	if err = Db.AutoMigrate(
		// migrating models
		&models.AuthorModel{},
		&models.BookModel{},
		&models.BookAvailabilityModel{},
		&models.BookGenreModel{},
		&models.BookRequestModel{},
		&models.BookRequestAcceptedModel{},
		&models.BookRequestDeclinedModel{},
		&models.BookRequestStatusModel{},
		&models.BookReturnsModel{},
		&models.ISBNVersionModel{},
	); err != nil {
		log.Fatal(err.Error())
	}
}

func SelectAllBooks() ([]models.BookModel, error) {
	books := []models.BookModel{}
	result := Db.Find(&books)
	return books, result.Error
}

func InsertBook(book *models.BookModel) (int, error) {
	result := Db.Create(book)
	return int(result.RowsAffected), result.Error
}

func SelectAllAvailableBooks() ([]models.BookModel, error) {
	books := []models.BookModel{}
	results := Db.Where("availability = available").Find(&books)

	return books, results.Error
}

func SelectBookById(id uuid.UUID) (models.BookModel, error) {
	book := models.BookModel{}
	result := Db.First(&book, id)
	return book, result.Error

}

func SelectBooksByName(name string) ([]models.BookModel, error) {
	books := []models.BookModel{}
	result := Db.Where("title = ?", name).Find(&books)
	return books, result.Error
}

func SelectAvailableBooksByName(name string) ([]models.BookModel, error) {
	books := []models.BookModel{}
	result := Db.Where("title = ? availability = available", name).Find(&books)
	return books, result.Error
}

func SelectAllRequests() ([]models.BookRequestModel, error) {
	reqs := []models.BookRequestModel{}
	result := Db.Find(&reqs)
	return reqs, result.Error
}

func InsertRequest(req *models.BookRequestModel) (int, error) {
	result := Db.Create(req)
	return int(result.RowsAffected), result.Error
}

func SelectAllAcceptedRequests() ([]models.BookRequestModel, error) {
	reqs := []models.BookRequestModel{}
	results := Db.Where("request_status = accepted").Find(&reqs)

	return reqs, results.Error
}

func SelectAllPendingRequests() ([]models.BookRequestModel, error) {
	reqs := []models.BookRequestModel{}
	results := Db.Where("request_status = pending").Find(&reqs)

	return reqs, results.Error
}

func SelectRequestById(id uuid.UUID) (models.BookRequestModel, error) {
	req := models.BookRequestModel{}
	result := Db.First(&req, id)
	return req, result.Error
}

func DeleteRequestById(id uuid.UUID) error {
	var req models.BookRequestModel
	req.Id = id
	result := Db.Delete(&req)
	return result.Error
}

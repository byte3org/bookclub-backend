package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Id              int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name            string `json:"name"`
	Email           string `json:"email"`
}

type BookGenre struct {
	gorm.Model
	Id              int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name            string `json:"name"`
}

type BookAvailability struct {
	gorm.Model
	Id           int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Availability string `json:"availability"`
}

type ISBNVersion struct {
	gorm.Model
	Id          int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name        string `json:"name"`
}

type Book struct {
	gorm.Model
	Id                      int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Title                   string `json:"title"`
	ISBN                    string `json:"isbn"`
	Picture                 string `json:"picture_url"`
	ISBNVersionId           int    `json:"isbn_version"`
    ISBNVersion             ISBNVersion `gorm:"foreignKey:ISBNVersionId"`
	AuthorId                int `json:"author"`
    Author                  Author `gorm:"foreignKey:AuthorId"`
	OwnerId                 int `json:"owner"`
	GenreId                 int `json:"genre"`
    BookGenre               BookGenre `gorm:"foreignKey:GenreId"`
	BookAvailabilityId      int `json:"availability"`
}

type BookRequestStatus struct {
    gorm.Model
	Id              int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Status          string `json:"status"`
}

type BookRequest struct {
    gorm.Model
	Id                      int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Title                   string `json:"title"`
	AuthorId                int    `json:"author"`
    Author                  Author `gorm:"foreignKey:AuthorId"`
	RequestorId             int `json:"requestor"`
	RequestStatusId         int `json:"request_status"`
    BookRequestStatus       BookRequestStatus `gorm:"foreignKey:RequestStatusId"`
	RequestedDate           time.Time `json:"requested_date" gorm:"datetime:timestamp;default:CURRENT_TIMESTAMP"`
}

type BookRequestAccepted struct {
    gorm.Model
	Id              int `json:"id" gorm:"primaryKey;autoIncrement:true"`
	BookId          int `json:"book"`
	BorrowerId      int `json:"borrower"`
	BorroweeId      int `json:"borrowee"`
	RequestId       int `json:"request"`
    BookRequest     BookRequest `gorm:"foreignKey:RequestId"`
	AcceptedDate time.Time `json:"accepted_date" gorm:"datetime:timestamp;default:CURRENT_TIMESTAMP"`
}

type BookRequestDeclined struct {
    gorm.Model
	Id              int `json:"id" gorm:"primaryKey;autoIncrement:true"`
	RequestId       int `json:"request"`
    BookRequest     BookRequest `gorm:"foreignKey:RequestId"`
	DeclinedUserId  int `json:"declined_user"`
}

type BookReturns struct {
    gorm.Model
	Id                  int `json:"id" gorm:"primaryKey;autoIncrement:true"`
	AcceptedRequestId   int `json:"accepted_request"`
    BookRequestAccepted BookRequestAccepted `gorm:"foreignKey:AcceptedRequestId"`
	ReturnedDate        time.Time `json:"returned_date" gorm:"datetime:timestamp;default:CURRENT_TIMESTAMP"`
}

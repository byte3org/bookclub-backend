package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserAddressInfo struct {
	gorm.Model
	Id         int     `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Address    string  `json:"address"`
	Longitudes float64 `json:"longitudes"`
	Latitudes  float64 `json:"latitudes"`
	UserId     int     `json:"user"`
	User       User
}

type UserContactInfo struct {
	gorm.Model
	Id      int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Contact string `json:"contact"`
	UserId  int    `json:"user"`
	User    User
}

type Author struct {
	gorm.Model
	Id    int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type BookGenre struct {
	gorm.Model
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string `json:"name"`
}

type BookAvailability struct {
	gorm.Model
	Id           int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Availability string `json:"availability"`
}

type ISBNVersion struct {
	gorm.Model
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string `json:"name"`
}

type Book struct {
	gorm.Model
	Id                 int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Title              string `json:"title"`
	ISBN               string `json:"isbn"`
	Picture            string `json:"picture_url"`
	ISBNVersionId      int    `json:"isbn_version"`
	ISBNVersion        ISBNVersion
	AuthorId           int `json:"author"`
	Author             Author
	OwnerId            int `json:"owner"`
	Owner              User
	GenreId            int `json:"genre"`
	Genre              BookGenre
	BookAvailabilityId int `json:"availability"`
	Availability       BookAvailability
}

type BookRequestStatus struct {
	Id     int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Status string `json:"status"`
}

type BookRequest struct {
	Id              int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Title           string `json:"title"`
	AuthorId        int    `json:"author"`
	Author          Author
	RequestorId     int `json:"requestor"`
	Requester       User
	RequestStatusId int `json:"request_status"`
	RequestStatus   BookRequestStatus
	RequestedDate   time.Time `json:"requested_date" gorm:"datetime:timestamp;default:CURRENT_TIMESTAMP"`
}

type BookRequestAccepted struct {
	Id           int `json:"id" gorm:"primaryKey;autoIncrement:true"`
	BookId       int `json:"book"`
	Book         Book
	BorrowerId   int `json:"borrower"`
	Borrower     User
	BorroweeId   int `json:"borrowee"`
	Borrowee     User
	RequestId    int `json:"request"`
	Request      BookRequest
	AcceptedDate time.Time `json:"accepted_date" gorm:"datetime:timestamp;default:CURRENT_TIMESTAMP"`
}

type BookRequestDeclined struct {
	Id             int `json:"id" gorm:"primaryKey;autoIncrement:true"`
	RequestId      int `json:"request"`
	Request        BookRequest
	DeclinedUserId int `json:"declined_user"`
	DeclinedUser   User
}

type BookReturns struct {
	Id                int `json:"id" gorm:"primaryKey;autoIncrement:true"`
	AcceptedRequestId int `json:"accepted_request"`
	AcceptedRequest   BookRequestAccepted
	ReturnedDate      time.Time `json:"returned_date" gorm:"datetime:timestamp;default:CURRENT_TIMESTAMP"`
}

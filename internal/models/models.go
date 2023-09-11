package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Id    int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name  string `json:"name"`
	Email string `json:"email"`
    Books        []Book `gorm:"foreignKey:AuthorId,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` 
    BookRequests []BookRequest `gorm:"foreignKey:AuthorId,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` 
}

type BookGenre struct {
	gorm.Model
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string `json:"name"`
    Books        []Book `gorm:"foreignKey:GenreId,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` 
}

type BookAvailability struct {
	gorm.Model
	Id           int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Availability string `json:"availability"`
    Books        []Book `gorm:"foreignKey:BookAvailabilityId,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` 
}

type ISBNVersion struct {
	gorm.Model
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string `json:"name"`
    Books        []Book `gorm:"foreignKey:ISBNVersionId,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` 
}

type Book struct {
	gorm.Model
	Id                 int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Title              string `json:"title"`
	ISBN               string `json:"isbn"`
	Picture            string `json:"picture_url"`
	ISBNVersionId      int    `json:"isbn_version"`
	AuthorId           int `json:"author"`
	OwnerId            int `json:"owner"`
	GenreId            int `json:"genre"`
	BookAvailabilityId int `json:"availability"`
    AcceptedBookRequests []BookRequestAccepted `gorm:"foreignKey:BookId,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`	
}

type BookRequestStatus struct {
    gorm.Model
	Id     int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Status string `json:"status"`
    BookRequests []BookRequest `gorm:"foreignKey:RequestStatusId,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type BookRequest struct {
    gorm.Model
	Id              int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Title           string `json:"title"`
	AuthorId        int    `json:"author"`
	RequestorId     int `json:"requestor"`	
	RequestStatusId int `json:"request_status"`
	RequestedDate   time.Time `json:"requested_date" gorm:"datetime:timestamp;default:CURRENT_TIMESTAMP"`
    AcceptedBookRequests BookRequestAccepted `gorm:"foreignKey:RequestId,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
    DeclinedBookRequests BookRequestDeclined `gorm:"foreignKey:RequestId,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type BookRequestAccepted struct {
    gorm.Model
	Id           int `json:"id" gorm:"primaryKey;autoIncrement:true"`
	BookId       int `json:"book"`
	BorrowerId   int `json:"borrower"`
	BorroweeId   int `json:"borrowee"`
	RequestId    int `json:"request"`
	AcceptedDate time.Time `json:"accepted_date" gorm:"datetime:timestamp;default:CURRENT_TIMESTAMP"`
    BookRequestReturn BookReturns `gorm:"foreignKey:AcceptedRequestId,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type BookRequestDeclined struct {
    gorm.Model
	Id             int `json:"id" gorm:"primaryKey;autoIncrement:true"`
	RequestId      int `json:"request"`
	DeclinedUserId int `json:"declined_user"`
}

type BookReturns struct {
    gorm.Model
	Id                int `json:"id" gorm:"primaryKey;autoIncrement:true"`
	AcceptedRequestId int `json:"accepted_request"`
	ReturnedDate      time.Time `json:"returned_date" gorm:"datetime:timestamp;default:CURRENT_TIMESTAMP"`
}

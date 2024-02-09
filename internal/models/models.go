package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	Id        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type AuthorModel struct {
	BaseModel
	Name  string `json:"name"`
	Email string `json:"email"`
}

type BookGenreModel struct {
	gorm.Model
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string `json:"name"`
}

type BookAvailabilityModel struct {
	gorm.Model
	Id           int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Availability string `json:"availability"`
}

type ISBNVersionModel struct {
	gorm.Model
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string `json:"name"`
}

type BookModel struct {
	BaseModel
	Title                 string                `json:"title"`
	ISBN                  string                `json:"isbn"`
	Picture               string                `json:"picture_url"`
	ISBNVersionId         int                   `json:"isbn_version"`
	ISBNVersion           ISBNVersionModel      `gorm:"foreignKey:ISBNVersionId"`
	AuthorId              uuid.UUID             `json:"author" gorm:"type:uuid"`
	Author                AuthorModel           `gorm:"foreignKey:AuthorId"`
	OwnerId               uuid.UUID             `json:"owner" gorm:"type:uuid"`
	GenreId               int                   `json:"genre"`
	BookGenre             BookGenreModel        `gorm:"foreignKey:GenreId"`
	BookAvailabilityId    int                   `json:"availability"`
	BookAvailabilityModel BookAvailabilityModel `gorm:"foreignKey:BookAvailabilityId"`
}

type BookRequestStatusModel struct {
	gorm.Model
	Id     int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Status string `json:"status"`
}

type BookRequestModel struct {
	BaseModel
	Title             string                 `json:"title"`
	Message           string                 `json:"message"`
	AuthorId          uuid.UUID              `json:"author" gorm:"type:uuid"`
	Author            AuthorModel            `gorm:"foreignKey:AuthorId"`
	RequestorId       uuid.UUID              `json:"requestor" gorm:"type:uuid"`
	RequestStatusId   int                    `json:"request_status"`
	BookRequestStatus BookRequestStatusModel `gorm:"foreignKey:RequestStatusId"`
	RequestedDate     time.Time              `json:"requested_date" gorm:"datetime:timestamp;default:CURRENT_TIMESTAMP"`
}

type BookRequestAcceptedModel struct {
	BaseModel
	BookId       uuid.UUID        `json:"book" gorm:"type:uuid"`
	Book         BookModel        `gorm:"foreignKey:BookId"`
	BorrowerId   uuid.UUID        `json:"borrower" gorm:"type:uuid"`
	BorroweeId   uuid.UUID        `json:"borrowee" gorm:"type:uuid"`
	RequestId    uuid.UUID        `json:"request" gorm:"type:uuid"`
	BookRequest  BookRequestModel `gorm:"foreignKey:RequestId"`
	AcceptedDate time.Time        `json:"accepted_date" gorm:"datetime:timestamp;default:CURRENT_TIMESTAMP"`
}

type BookRequestDeclinedModel struct {
	BaseModel
	RequestId      uuid.UUID        `json:"request" gorm:"type:uuid"`
	BookRequest    BookRequestModel `gorm:"foreignKey:RequestId"`
	DeclinedUserId uuid.UUID        `json:"declined_user" gorm:"type:uuid"`
}

type BookReturnsModel struct {
	BaseModel
	AcceptedRequestId   uuid.UUID                `json:"accepted_request" gorm:"type:uuid"`
	BookRequestAccepted BookRequestAcceptedModel `gorm:"foreignKey:AcceptedRequestId"`
	ReturnedDate        time.Time                `json:"returned_date" gorm:"datetime:timestamp;default:CURRENT_TIMESTAMP"`
}

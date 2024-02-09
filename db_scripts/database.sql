CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE 
    IF NOT EXISTS Authors (
        id UUID PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()), 
        name VARCHAR(20) NOT NULL,
        email VARCHAR(20)
    );

CREATE TABLE 
    IF NOT EXISTS BookGenres (
        id INT PRIMARY KEY NOT NULL GENERATED ALWAYS AS IDENTITY,
        name VARCHAR(10) NOT NULL
    );

/*
available
borrowed
lost
*/
CREATE TABLE 
    IF NOT EXISTS BookAvailability (
        id INT PRIMARY KEY NOT NULL GENERATED ALWAYS AS IDENTITY,
        availability VARCHAR(20) NOT NULL
    );

CREATE TABLE 
    IF NOT EXISTS ISBNVersion (
        id INT PRIMARY KEY NOT NULL GENERATED ALWAYS AS IDENTITY,
        name VARCHAR(7) NOT NULL
    );

CREATE TABLE 
    IF NOT EXISTS Books (
        id UUID PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),      
        title VARCHAR(20) NOT NULL,
        isbn VARCHAR(20) NOT NULL,
        picture_url VARCHAR(100) NOT NULL, 
        isbn_version INT NOT NULL,
        CONSTRAINT fk_isbnversion FOREIGN KEY(isbn_version) REFERENCES ISBNVersion(id),
        author UUID NOT NULL,
        CONSTRAINT fk_author FOREIGN KEY(author) REFERENCES Authors(id),
        owner UUID NOT NULL,
        genre INT NOT NULL,
        CONSTRAINT fk_genre FOREIGN KEY(genre) REFERENCES BookGenres(id),
        availability INT NOT NULL,
        CONSTRAINT fk_availability FOREIGN KEY(availability) REFERENCES BookAvailability(id)
    );

/*
pending
accepted
returned
*/
CREATE TABLE 
    IF NOT EXISTS BookRequestStatus (
        id INT PRIMARY KEY NOT NULL GENERATED ALWAYS AS IDENTITY,
        status VARCHAR(10) NOT NULL
    );

/* 
these are public book requests, meaning that every owner within that raduis gets a notify
and can accept or reject

personal ones on the hand are sent to specific book owners and for specific book ID
 */
CREATE TABLE 
    IF NOT EXISTS BookRequests (
        id UUID PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),    
        title VARCHAR(20) NOT NULL,
        message VARCHAR(255) NOT NULL,
        author UUID NOT NULL,
        CONSTRAINT fk_author FOREIGN KEY(author) REFERENCES Authors(id),
        requestor UUID NOT NULL, 
        request_status INT NOT NULL,
        CONSTRAINT fk_status FOREIGN KEY(request_status) REFERENCES BookRequestStatus(id),
        requested_date TIMESTAMP
    );

CREATE TABLE 
    IF NOT EXISTS BookRequestAccepted (
        id UUID PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),   
        book UUID NOT NULL,
        CONSTRAINT fk_book FOREIGN KEY(book) REFERENCES Books(id),
        borrower UUID NOT NULL,
        borrowee UUID NOT NULL, 
        request UUID NOT NULL,
        CONSTRAINT fk_request FOREIGN KEY(request) REFERENCES BookRequests(id),
        accepted_date TIMESTAMP
    );

CREATE TABLE 
    IF NOT EXISTS BookRequestDeclined (
        id UUID PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),   
        request UUID NOT NULL,
        CONSTRAINT fk_request FOREIGN KEY(request) REFERENCES BookRequests(id),
        declined_user UUID NOT NULL
    );

CREATE TABLE 
    IF NOT EXISTS BookReturns (
        id UUID PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),   
        accepted_request UUID NOT NULL,
        CONSTRAINT fk_accepted_request FOREIGN KEY(accepted_request) REFERENCES BookRequestAccepted(id),
        returned_date TIMESTAMP
    );

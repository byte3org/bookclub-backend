CREATE TABLE Users {
    id INT GENERATED ALWAYS AS IDENTITY,
    username UNIQUE VARCHAR(20),
    email UNIQUE VARCHAR(30),
    password VARCHAR(20),
    PRIMARY KEY(id)
};

CREATE TABLE UserAddressInfo {
    id INT GENERATED ALWAYS AS IDENTITY,
    address VARCHAR(50),
    longitudes FLOAT,
    latitudes FLOAT,
    CONSTRAINT fk_user FOREIGN KEY(user) REFERENCES Users(id),
    PRIMARY KEY(id)
};

CREATE TABLE UserContactInfo {
    id INT GENERATED ALWAYS AS IDENTITY,
    contact VARCHAR(10)
    CONSTRAINT fk_user FOREIGN KEY(user) REFERENCES Users(id),
    PRIMARY KEY(id)
};

CREATE TABLE Authors {
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(20),
    email VARCHAR(20),
};

CREATE TABLE BookGenres {
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(10),
};

/*
available
borrowed
lost
*/
CREATE TABLE BookAvailability {
    id INT GENERATED ALWAYS AS IDENTITY,
    availability VARCHAR(20),
    PRIMARY KEY(id)
};

CREATE TABLE ISBNVersion {
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(7),
    PRIMARY KEY(id)
}

CREATE TABLE Books {
    id INT GENERATED ALWAYS AS IDENTITY,
    title VARCHAR(20),
    isbn VARCHAR(20),
    picture_url VARCHAR(100),
    CONSTRAINT fk_isbnversion FOREIGN KEY(isbn_version) REFERENCES ISBNVersion(id),
    CONSTRAINT fk_author FOREIGN KEY(author) REFERENCES Authors(id),
    CONSTRAINT fk_owner FOREIGN KEY(owner) REFERENCES Users(id),
    CONSTRAINT fk_genre FOREIGN KEY(genre) REFERENCES BookGenres(id),
    CONSTRAINT fk_availability FOREIGN KEY(availability) REFERENCES BookAvailability(id).
    PRIMARY KEY(id)
};

/*
pending
accepted
returned
*/
CREATE TABLE BookRequestStatus {
    id INT GENERATED ALWAYS AS IDENTITY,
    status VARCHAR(10),
    PRIMARY KEY(id)
};

CREATE TABLE BookRequests {
    id INT GENERATED ALWAYS AS IDENTITY,
    title VARCHAR(20),
    CONSTRAINT fk_author FOREIGN KEY(author) REFERENCES Authors(id),
    CONSTRAINT fk_requestor FOREIGN KEY(requestor) REFERENCES Users(id),
    CONSTRAINT fk_status FOREIGN KEY(request_status) REFERENCES BookRequestStatus(id),
    requested_date TIMESTAMP,
    PRIMARY KEY(id)
};

CREATE TABLE BookRequestAccepted {
    id INT GENERATED ALWAYS AS IDENTITY,
    CONSTRAINT fk_book FOREIGN KEY(book) REFERENCES Books(id),
    CONSTRAINT fk_borrower FOREIGN KEY(borrower) REFERENCES Users(id),
    CONSTRAINT fk_borrowee FOREIGN KEY(borrowee) REFERENCES Users(id),
    CONSTRAINT fk_request FOREIGN KEY(request) REFERENCES BookRequests(id),
    accepted_date TIMESTAMP,
    PRIMARY KEY(id),
};

CREATE TABLE BookRequestDeclined {
    id INT GENERATED ALWAYS AS IDENTITY,
    CONSTRAINT fk_request FOREIGN KEY(request) REFERENCES BookRequests(id),
    CONSTRAINT fk_declined_user FOREIGN KEY(declined_user) REFERENCES Users(id),
    PRIMARY KEY(id)
}

CREATE TABLE BookReturns {
    id INT GENERATED ALWAYS AS IDENTITY,
    CONSTRAINT fk_accepted_request FOREIGN KEY(accepted_request) BookRequestAccepted(id),
    returned_date TIMESTAMP,
    PRIMARY KEY(id)
}


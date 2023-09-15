FROM golang:1.20

WORKDIR /bookclub-backend

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

COPY . .

EXPOSE 8001

CMD ["bash","/wait-for-it.sh", "db:5432", "--", "go", "run", "./cmd/backend/main.go"]

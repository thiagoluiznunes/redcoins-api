FROM golang:latest

LABEL maintainer="thiagoluiz.dev@gmail.com"

# RUN apt-get update
# RUN apt-get install vim -y
# RUN apt-get install iputils-ping
RUN export GOBIN=$GOPATH/bin

WORKDIR /go/src/redcoins-api

COPY . ./

RUN go get github.com/joho/godotenv
RUN go get github.com/go-chi/chi
RUN go get github.com/go-chi/render
RUN go get github.com/go-chi/cors
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/dgrijalva/jwt-go
RUN go get golang.org/x/crypto/bcrypt

RUN go build -o bin/main cmd/redcoins-api/main.go

CMD ["./bin/main"]

EXPOSE 8000/tcp

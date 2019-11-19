package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

// DB : database instance
var DB *sql.DB

// User : user model
type User struct {
	uuid     string
	name     string
	email    string
	password string
}

// InitUserSchema : init table
func InitUserSchema() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS Users (
		uuid varchar(36) NOT NULL,
		name varchar(100) NOT NULL,
		email varchar(100) NOT NULL,
		password varchar(100) NOT NULL,
		PRIMARY KEY (uuid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`)

	if err != nil {
		panic(err)
	}
}

// CreateUserService : create user in database
func CreateUserService(user User) (string, error) {
	var email string
	selectUserQuery := fmt.Sprintf("SELECT email FROM Users WHERE email = '%s'", user.email)
	row := DB.QueryRow(selectUserQuery)
	err := row.Scan(&email)

	if err == sql.ErrNoRows {
		insertUserQuery := fmt.Sprintf(`
		INSERT INTO Users (uuid, name, email, password)
		VALUES (UUID(), '%s',	'%s',	'%s')`, user.name, user.email, user.password)
		insert, err := DB.Query(insertUserQuery)
		insert.Close()

		if err != nil {
			panic(err.Error())
		}
		return "Token", nil
	} else if err != nil {
		log.Fatal(err)
		return "", err
	}

	if user.email != "" {
		return "", errors.New("user: Already registered")
	}

	return "", errors.New("user: Not found")
}

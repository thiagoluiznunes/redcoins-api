package user

import (
	"database/sql"
	"fmt"
	"log"
)

// DB : database instance
var DB *sql.DB

// User : user model
type User struct {
	name     string
	email    string
	password string
}

// InitUserSchema : init table
func InitUserSchema() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS Users (
		uuid BINARY(16) NOT NULL,
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
func CreateUserService(user User) (*User, error) {
	var s sql.NullString
	findUserQuery := fmt.Sprintf("SELECT * FROM Users WHERE email = '%s'", user.email)
	err := DB.QueryRow(findUserQuery).Scan(&s)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("User Not found.")
		// createuserQuery := fmt.Sprintf(`
		// 	INSERT INTO Users (uuid, name, email, password)
		// 	VALUES (UUID_TO_BIN(UUID()), %s,	%s,	%s`, user.name, user.email, user.password)

	case err != nil:
		log.Fatal(err)
	default:
		// do stuffs
	}
	return nil, nil
}

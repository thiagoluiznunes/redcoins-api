package user

import (
	"database/sql"
	"errors"
	"fmt"
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
		CREATE TABLE IF NOT EXISTS users (
		uuid VARCHAR(36) NOT NULL UNIQUE,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(100) NOT NULL,
		CONSTRAINT pk_user_uuid PRIMARY KEY (uuid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`)

	if err != nil {
		panic(err)
	}
}

// CreateUser : create user in database
func CreateUser(user User) error {
	var email string
	selectUserQuery := fmt.Sprintf("SELECT email FROM users WHERE email = '%s'", user.email)
	row := DB.QueryRow(selectUserQuery)
	err := row.Scan(&email)

	if err == sql.ErrNoRows {
		insertUserQuery := fmt.Sprintf(`
			INSERT INTO users (uuid, name, email, password)
			VALUES (UUID(), '%s', '%s', '%s');`, user.name, user.email, user.password)

		insert, err := DB.Query(insertUserQuery)
		insert.Close()

		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}

	if user.email != "" {
		return errors.New("user: Already registered")
	}
	return errors.New("user: Not found")
}

// FindUserByEmail : find user by email in database
func FindUserByEmail(email string) (User, error) {
	var user User

	selectUserQuery := fmt.Sprintf(`SELECT * FROM users	WHERE email = '%s'`, email)

	row := DB.QueryRow(selectUserQuery)
	err := row.Scan(&user.uuid, &user.name, &user.email, &user.password)

	if err == sql.ErrNoRows {
		return user, errors.New("user: User not found")
	} else if err != nil {
		return user, err
	} else if user.uuid != "" {
		return user, nil
	}
	return user, errors.New("user: Not found")
}

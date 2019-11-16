package user

import (
	"database/sql"
	"fmt"
)

func init() {
	fmt.Println("Init model.")
}

// User : user model
type User struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

// AllUsers : retrive all user from database
func AllUsers(db *sql.DB) ([]*User, error) {
	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*User, 0)
	for rows.Next() {
		bk := new(User)
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}

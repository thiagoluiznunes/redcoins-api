package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// DataBase : database structure
// type DataBase struct {
// 	conn *sql.DB
// }

// InitDataBase : database structure
func InitDataBase() *sql.DB {
	fmt.Println("")
	dbHost := os.Getenv("DB_HOST")
	dbDialect := os.Getenv("DB_DIALECT")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open(dbDialect, url)

	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return db
	// return &DataBase{conn: db}
}

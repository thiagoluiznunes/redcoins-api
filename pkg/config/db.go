package main

import (
	"database/sql"
  _"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func InitConnection() {
	db, err = sql.Open("mysql", "redventures:redventures123@tcp(127.0.0.1:3306)/redcoins")
  if err != nil {
    panic(err.Error())
	}
	defer db.Close()
}

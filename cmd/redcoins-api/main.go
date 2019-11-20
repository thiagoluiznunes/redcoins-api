package main

import (
	"log"
	"net/http"
	"redcoins-api/configs"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := configs.InitDataBase()
	router := configs.InitRoutes(db)
	http.ListenAndServe(":8000", router)
	log.Fatal(http.ListenAndServe(":8000", router))
}

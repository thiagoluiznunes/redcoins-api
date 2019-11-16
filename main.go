package main

import (
	"net/http"
	"rv-api/configs"
)

func main() {
	router := configs.InitRoutes()
	http.ListenAndServe(":8000", router)
}

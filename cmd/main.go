package main

import (
	"github.com/timurkaev/blog-backend/database"
	"github.com/timurkaev/blog-backend/routes"
	"log"
	"net/http"
)

func main() {
	database.ConnectDb()
	router := routes.SetupRouter()

	log.Fatal(http.ListenAndServe(":3000", router))
}

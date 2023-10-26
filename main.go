package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/janapc/game-api-go/database"
	"github.com/janapc/game-api-go/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.Connection()
	defer database.Database.Close()
	http.HandleFunc("/game", routes.RouteGame)
	http.HandleFunc("/game/", routes.RouteGameByParam)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

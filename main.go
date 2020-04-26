package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aswinudhayakumar/golang-boilerplate/controller"
	"github.com/aswinudhayakumar/golang-boilerplate/driver"
	"github.com/aswinudhayakumar/golang-boilerplate/middleware"
	"github.com/aswinudhayakumar/golang-boilerplate/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Cannot find .env file")
		return
	}

	serverPort := os.Getenv("SERVER_PORT")
	serverPort = ":" + serverPort

	if serverPort == ":" {
		fmt.Println("Unable to find SERVER_PORT from environmental variables")
		return
	}

	db, err := driver.Connect()

	userController := controller.UserController{DB: db}

	if err == nil {
		router := mux.NewRouter().StrictSlash(true)
		router.Use(middleware.AuthenticationMiddleware)
		routes.HandleUserRoutes(userController, router)
		fmt.Println("GO server started and running at port 8123")
		log.Fatal(http.ListenAndServe(serverPort, cors.AllowAll().Handler(router)))
		fmt.Println("Server stopped")
	}
}

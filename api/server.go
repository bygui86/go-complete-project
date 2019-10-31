package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/bygui86/go-complete-project/api/controllers"
	"github.com/bygui86/go-complete-project/api/seed"
)

var server = controllers.Server{}

func Run() {
	// load env vars
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	// init server
	server.Init(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	// seed DB
	seed.Load(server.DB)
	// run server
	server.Run(":8080")
}

package main

import (
	"library/config"
	"library/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// Connect to MongoDB (NEW)
	config.ConnectMongo()

	r := routes.SetUpRouter()
	r.Run(":8000")
}

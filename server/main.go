package main

import (
	"library/config"
	"library/routes"
	"library/genai_client"

	"github.com/joho/godotenv"
	"log"
)

func main() {
	godotenv.Load()

	config.ConnectMongo()

	if err := genai_client.InitClient(); err != nil {
		log.Fatal("❌ Gemini client initialization failed:", err)
	}

	r := routes.SetUpRouter()
	r.Run(":8000")
}

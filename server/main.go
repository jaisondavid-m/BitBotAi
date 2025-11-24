package main

import (
	"library/config"
	"library/routes"

	"github.com/joho/godotenv"
)
func main(){
	godotenv.Load()
	config.Connect()
	defer config.DB.Close()
	r:=routes.SetUpRouter()
	r.Run(":8000")
}
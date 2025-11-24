package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
func Connect(){
	var err error
	DB,err = sql.Open("mysql","root:MySqlJaison007@@tcp(localhost:3306)/bitbotai")
	if err!=nil{
		log.Fatal(err)
	}
	if err=DB.Ping();err !=nil{
		log.Fatal("Failed to connect to database",err)
	}
	log.Println("DB connected successfully")
}
package main

import (
	"library/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"
)
func main(){
	godotenv.Load()

	r:=gin.Default()
	r.Use(cors.Default())

	r.Use(func(c *gin.Context){
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS"{
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/upload",handlers.UploadText)
	r.POST("/ask",handlers.AskQuestions)
	fmt.Println("server runnning on port 8000")
	r.Run(":8000")
}
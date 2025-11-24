package routes

import (
	"library/handlers"
	"library/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine{
	r:=gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
	}))

	r.POST("/login",handlers.Login)
	r.POST("/register",handlers.Register)
	r.POST("/ask",middlewares.Protected(),handlers.AskQuestions)
	r.POST("/upload",middlewares.Protected(),middlewares.IsAdmin(),handlers.UploadText)
	return r
}
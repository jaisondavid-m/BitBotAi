package middlewares

import (
	"library/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey=[]byte("It_is_not_a_secret_key_but_a_secret_key")

func Protected() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenstring := c.GetHeader("Authorization")

		if tokenstring==""{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"No Token Proveide"})
			return
		}

		token,err:=jwt.Parse(tokenstring,func (token *jwt.Token)(interface{},error){
			return jwtKey,nil
		})

		if err!=nil || !token.Valid{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid or Expired Token"})
			c.Abort()
			return 
		}
		c.Next()
	}
}
func IsAdmin() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenstring := c.GetHeader("Authorization")

		token,err:=jwt.ParseWithClaims(tokenstring,&models.Claims{},func(token *jwt.Token)(interface{},error){
			return jwtKey,nil
		})

		if err!=nil || !token.Valid{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid or Expired Token"})
			c.Abort()
			return 
		}

		claims,ok:=token.Claims.(*models.Claims)
		if !ok{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Failed To Read Token"})
			return 
		}

		if claims.Role != "admin"{
			c.JSON(http.StatusForbidden,gin.H{"error":"Access Denied"})
			c.Abort()
			return 
		}

		c.Next()
	}
}
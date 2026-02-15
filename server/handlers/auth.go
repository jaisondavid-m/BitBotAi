package handlers

import (
	"library/storage"
	"library/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func Login(c *gin.Context){
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := storage.FindUserByEmail(input.Email)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid Email"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid Password"})
		return
	}

	expirationTime := time.Now().Add(72*time.Hour)
	claims := &models.Claims{
		Name:user.Name,
		Email:user.Email,
		Role:user.Role,
		RegisteredClaims:jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenstring,err:=token.SignedString(jwtKey)
	if err !=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Could not generate Token"})
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"Logged In Successfully",
		"token":tokenstring,
	})
}

func Register(c *gin.Context) {
	var input models.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// check email exists
	_, err := storage.FindUserByEmail(input.Email)
	if err == nil {
		c.JSON(400, gin.H{"error": "User Already Exists"})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Role:     "user",
		Password: string(hash),
	}

	if err := storage.CreateUser(user); err != nil {
		c.JSON(500, gin.H{"error": "Failed To Register"})
		return
	}

	c.JSON(200, gin.H{"message": "Registered Successfully"})
}

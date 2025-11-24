package handlers

import (
	"library/config"
	"library/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey=[]byte("It_is_not_a_secret_key_but_a_secret_key")

func Login(c *gin.Context){
	var input models.LoginInput
	if err:=c.ShouldBindJSON(&input); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	var user models.User
	err:=config.DB.QueryRow("SELECT name,email,role,password FROM users WHERE email=(?)",input.Email).Scan(&user.Name,&user.Email,&user.Role,&user.Password);

	if err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid Email"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(input.Password))
	if err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid Password"})
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

func Register(c *gin.Context){
	var input models.RegisterInput
	if err:= c.ShouldBindJSON(&input); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	
	var existingUser string
	err := config.DB.QueryRow("SELECT email from users where email = (?)",input.Email).Scan(&existingUser)
	if err==nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"User Already Exists"})
		return
	}

	hashedpassword,err := bcrypt.GenerateFromPassword([]byte(input.Password),bcrypt.DefaultCost)

	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Server Error"})
		return
	}

	_,err = config.DB.Exec("INSERT INTO users (name,email,role,password) VALUES (?,?,?,?)",input.Name,input.Email,"user",string(hashedpassword))
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed To Register"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"messsage":"Registered Successfully"})
}
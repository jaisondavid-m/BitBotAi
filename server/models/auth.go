package models

import "github.com/golang-jwt/jwt/v5"


type User struct{
	Name			string		`json:"name"`
	Email			string		`json:"email"`
	Password		string		`json:"password"`
	Role			string		`json:"role"`
}
type RegisterInput struct{
	Name			string		`json:"name"`
	Email			string		`json:"email"`
	Password		string		`json:"password"`
	Role			string		`json:"role"`	
}
type LoginInput struct{
	Email			string		`json:"email"`
	Password		string		`json:"password"`
}
type Claims struct{
	Name			string		`json:"name"`
	Email			string		`json:"email"`
	Role			string		`json:"role"`	
	jwt.RegisteredClaims
}
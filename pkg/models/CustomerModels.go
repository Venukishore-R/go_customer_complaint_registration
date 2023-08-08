package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Customer struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type CustomerRepository struct {
	customers []*Customer
}

var TokenString string

type Claims struct {
	Username string
	Email    string
	Phone    string
	jwt.StandardClaims
}

var Jwtkey = []byte("secret_key")

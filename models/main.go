package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/dgrijalva/jwt-go"
)

type User struct {
  gorm.Model
  Id string
  Email string
  Password string
  Name string
}

type Record struct {
  gorm.Model
  Id string
  Email string
  Password string
  Name string
}

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Login bool   `json:"login"`
	jwt.StandardClaims
}

package config

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/dgrijalva/jwt-go"
)

type User struct {
  gorm.Model
  Email     string  `gorm:"type:varchar(100);unique_index;not null" form:"email"`
  Password  string  `gorm:"type:varchar(32);not null" form:"password"`
  Name      string  `gorm:"type:varchar(64);not null" form:"name"`
}

type Record struct {
  gorm.Model
  Title       string  `gorm:"type:varchar(100)" form:"title"`
  Content     string  `gorm:"type:varchar(2048)" form:"content"`
  Type        string  `gorm:"type:varchar(20)"`
  IconUrl     string  `gorm:"type:varchar(256)" form:"icon_url"`
  CreatorId   uint    `gorm:"FOREIGNKEY;not null"`
}

type JwtCustomClaims struct {
  ID    uint   `json:"id"`
  Name  string `json:"name"`
  jwt.StandardClaims
}

type PageableParam struct {
  // Page  uint `json:"page" form:"page" query:"page" validate:"min=1"`
  Page  uint `query:"page" validate:"min=1"`
  Limit uint `query:"limit" validate:"min=1"`
}

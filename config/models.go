package config

import (
  "time"

  "github.com/dgrijalva/jwt-go"
)

type (
  User struct {
    ID        uint `gorm:"primary_key" json:"id"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index"`

    Name      string  `gorm:"type:varchar(64);not null" form:"name" json:"name"`
    Email     string  `gorm:"type:varchar(100);unique;index;not null" form:"email" json:"email"`
    Password  string  `gorm:"type:varchar(64);not null" form:"password"`
  }

  Record struct {
    ID        uint `gorm:"primary_key" json:"id"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index"`

    Title     string  `gorm:"type:varchar(100)" form:"title" json:"title"`
    Content   string  `gorm:"type:varchar(2048)" form:"content" json:"content"`
    Type      string  `gorm:"type:varchar(20)" json:"type"`
    IconUrl   string  `gorm:"type:varchar(256)" form:"icon_url" json:"icon_url"`
    CreatorId uint    `gorm:"FOREIGNKEY;not null"`
  }

  JwtCustomClaims struct {
    ID    uint   `json:"id"`
    Email string `json:"email"`
    jwt.StandardClaims
  }
)

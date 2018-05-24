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

    Name      string  `gorm:"type:varchar(64);not null" json:"name"`
    Email     string  `gorm:"type:varchar(100);unique;index;not null" json:"email"`
    AvatarUrl string  `gorm:"type:varchar(256)" json:"icon_url"`
    Password  string  `gorm:"type:varchar(64);not null"`
  }

  Record struct {
    ID        uint `gorm:"primary_key" json:"id"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index"`

    Title     string  `gorm:"type:varchar(100)" json:"title"`
    Content   string  `gorm:"type:varchar(2048)" json:"content"`
    Type      string  `gorm:"type:varchar(20)" json:"type"`
    IconUrl   string  `gorm:"type:varchar(256)" json:"icon_url"`
    CreatorId uint    `gorm:"FOREIGNKEY;not null"`
  }

  JwtCustomClaims struct {
    ID    uint   `json:"id"`
    Email string `json:"email"`
    jwt.StandardClaims
  }
)

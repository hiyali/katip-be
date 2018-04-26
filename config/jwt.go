package config

import (
  "github.com/labstack/echo/middleware"
)

func GetJwtConfig () middleware.JWTConfig {
  return middleware.JWTConfig{
    Claims:     &JwtCustomClaims{},
    SigningKey: []byte("katip_known_secret"),
  }
}

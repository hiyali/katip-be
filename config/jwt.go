package config

import "github.com/labstack/echo/middleware"

func GetJwtConfig () middleware.JWTConfig {
  cf := Get()
  return middleware.JWTConfig{
    Claims:     &JwtCustomClaims{},
    SigningKey: []byte(cf.App.Secret),
  }
}

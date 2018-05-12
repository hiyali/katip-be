package handlers

import (
  "net/http"

  "github.com/labstack/echo"
  "github.com/dgrijalva/jwt-go"

  "github.com/hiyali/katip-be/config"
)

func UserGetInfo(c echo.Context) (err error) {
  user := c.Get("user").(*jwt.Token)
  claims := user.Claims.(*config.JwtCustomClaims)

  return c.JSON(http.StatusOK, config.JsonUser{
    claims.ID,
    claims.Name,
    claims.Email,
  })
}

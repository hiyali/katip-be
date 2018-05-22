package handlers

import (
  "net/http"

  "github.com/labstack/echo"
  _ "github.com/jinzhu/gorm/dialects/mysql"

  "github.com/hiyali/katip-be/config"
)

func Ping(c echo.Context) error {
  cf := config.Get()
  return c.String(http.StatusOK, cf.App.Name + " - api ping")
}

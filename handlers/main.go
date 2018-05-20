package handlers

import (
  "net/http"

  "github.com/labstack/echo"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Ping(c echo.Context) error {
  return c.String(http.StatusOK, "Katip - api ping")
}

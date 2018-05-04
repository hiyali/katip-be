package handlers

import (
  "net/http"

	"github.com/labstack/echo"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Katip - home")
}

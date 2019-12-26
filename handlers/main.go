package handlers

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"

	"github.com/hiyali/katip-be/config"
)

func Ping(c echo.Context) error {
	cf := config.Get()
	return c.String(http.StatusOK, cf.App.Name+" - api ping")
}

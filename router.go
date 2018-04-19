package main

import (
	"github.com/labstack/echo"

  "github.com/hiyali/katip-be/handlers"
)

func RegisterRouters (e *echo.Echo) {
  handlers.RegisterRouters(e)
}

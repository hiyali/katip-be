package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

  routerRegister(e)

	// Start server
	e.Logger.Fatal(e.Start(":5555"))
}

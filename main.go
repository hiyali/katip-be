package main

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

func main() {
  e := echo.New()

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.Validator = &Validator{}

  routerRegister(e)

  // Start server
  e.Logger.Fatal(e.Start(":5555"))
}

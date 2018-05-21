package main

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

func main() {
  e := echo.New()

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    Skipper:      middleware.DefaultSkipper,
    AllowOrigins: []string{"*"},
    AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
  }))

  e.Validator = &Validator{}

  routerRegister(e)

  // Start server
  e.Logger.Fatal(e.Start(":5555"))
}

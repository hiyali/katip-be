package main

import (
	"github.com/labstack/echo"
  "github.com/labstack/echo/middleware"

  "github.com/hiyali/katip-be/config"
  "github.com/hiyali/katip-be/handlers"
)

func routerRegister (e *echo.Echo) {
  // Login route
  e.POST("/login", handlers.UserLogin)

  // Unauthenticated route
  e.GET("/", handlers.Accessible)

  // Restricted group
  r := e.Group("/restricted")

  // Configure middleware with the custom claims type
  jwtConfig := config.GetJwtConfig()
  r.Use(middleware.JWTWithConfig(jwtConfig))
  r.GET("", handlers.UserRestricted)
}

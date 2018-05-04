package main

import (
	"github.com/labstack/echo"
  "github.com/labstack/echo/middleware"

  "github.com/hiyali/katip-be/config"
  "github.com/hiyali/katip-be/handlers"
)

func routerRegister (e *echo.Echo) {
  // Home
  e.GET("/", handlers.Home)

  // Auth
  userGroup := e.Group("/user")
  userGroup.POST("/login", handlers.UserLogin)
  // userGroup.POST("logout", handlers.UserLogout) // server side logout

  // Record (need login)
  recordGroup := e.Group("/record")
  jwtConfig := config.GetJwtConfig()
  recordGroup.Use(middleware.JWTWithConfig(jwtConfig))
  recordGroup.GET("", handlers.RecordGetAllPageable)
  recordGroup.POST("", handlers.RecordCreateOne)
  recordGroup.GET("/:id", handlers.RecordGetOne)
  recordGroup.PUT("/:id", handlers.RecordUpdateOne)
  recordGroup.DELETE("/:id", handlers.RecordDeleteOne)
}

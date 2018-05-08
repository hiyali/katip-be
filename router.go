package main

import (
	"github.com/labstack/echo"
  "github.com/labstack/echo/middleware"

  "github.com/hiyali/katip-be/config"
  "github.com/hiyali/katip-be/handlers"
)

func routerRegister (e *echo.Echo) {
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    Skipper:      middleware.DefaultSkipper,
    AllowOrigins: []string{"*"},
    AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
  }))

  // Home
  e.GET("/", handlers.Home)

  // Auth
  e.POST("/auth/login", handlers.AuthLogin)

  // JWT config
  jwtConfig := config.GetJwtConfig()

  // User Group (need login)
  ug := e.Group("/user")
  ug.Use(middleware.JWTWithConfig(jwtConfig))
  ug.GET("", handlers.UserGetInfo)
  // ug.PUT("", handlers.UserUpdateInfo)
  // ug.POST("", handlers.UserCreate)
  // ug.POST("/logout", handlers.UserLogout) // server side logout

  // Record Group (need login)
  rg := e.Group("/record")
  rg.Use(middleware.JWTWithConfig(jwtConfig))
  rg.GET("", handlers.RecordGetAllPageable)
  rg.POST("", handlers.RecordCreateOne)
  rg.GET("/:id", handlers.RecordGetOne)
  rg.PUT("/:id", handlers.RecordUpdateOne)
  rg.DELETE("/:id", handlers.RecordDeleteOne)
}

package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/hiyali/katip-be/config"
	"github.com/hiyali/katip-be/handlers"
)

func routerRegister(e *echo.Echo) {
	// JWT config
	jwtConfig := config.GetJwtConfig()

	// Ping
	e.GET("/api/ping", handlers.Ping)

	// Auth
	e.POST("/api/auth/login", handlers.AuthLogin)

	// User Register
	e.POST("/api/register", handlers.UserRegister)
	e.GET("/api/register-confirm", handlers.UserRegisterConfirm)
	e.POST("/api/forget-password", handlers.UserForgetPassword)
	e.POST("/api/reset-password", handlers.UserResetPassword)
	// User Group (need login)
	ug := e.Group("/api/user")
	ug.Use(middleware.JWTWithConfig(jwtConfig))
	ug.GET("", handlers.UserGetInfo)
	ug.PUT("", handlers.UserUpdateInfo)
	ug.PUT("/change-password", handlers.UserChangePassword)
	// ug.POST("/logout", handlers.UserLogout) // server side logout

	// Record Group (need login)
	rg := e.Group("/api/record")
	rg.Use(middleware.JWTWithConfig(jwtConfig))
	rg.GET("", handlers.RecordGetAllPageable)
	rg.POST("", handlers.RecordCreateOne)
	rg.GET("/:id", handlers.RecordGetOne)
	rg.PUT("/:id", handlers.RecordUpdateOne)
	rg.DELETE("/:id", handlers.RecordDeleteOne)
}

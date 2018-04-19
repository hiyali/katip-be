package handlers

import (
  "net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

  "github.com/hiyali/katip-be/models"
)

func RegisterRouters (e *echo.Echo) {
  // Login route
  e.POST("/login", userLogin)

  // Unauthenticated route
  e.GET("/", accessible)

  // Restricted group
  r := e.Group("/restricted")

  // Configure middleware with the custom claims type
  config := middleware.JWTConfig{
    Claims:     &models.JwtCustomClaims{},
    SigningKey: []byte("katip_known_secret"),
  }
  r.Use(middleware.JWTWithConfig(config))
  r.GET("", userRestricted)
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

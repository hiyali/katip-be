package handlers

import (
  "time"
  "net/http"

	"github.com/labstack/echo"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/dgrijalva/jwt-go"

  "github.com/hiyali/katip-be/database"
  "github.com/hiyali/katip-be/models"
)

// Login with email & password
func userLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

  db, err := database.GetDB()
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()

  var user models.User
  db.Where("email = ?", email).First(&user)

	if password == user.Password {
		// Set custom claims
		claims := &models.JwtCustomClaims{
			user.Name,
			true,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("katip_known_secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

// Authorization: Bearer {TOKEN_HERE}
func userRestricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

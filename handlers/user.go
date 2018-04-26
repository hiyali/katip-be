package handlers

import (
  "time"
  "net/http"

	"github.com/labstack/echo"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/dgrijalva/jwt-go"

  "github.com/hiyali/katip-be/config"
)

// Login with email & password
func UserLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

  db, err := config.GetDB()
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()

  var user config.User
  db.Where("email = ?", email).First(&user)

	if password == user.Password {
		// Set custom claims
		claims := &config.JwtCustomClaims{
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
func UserRestricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

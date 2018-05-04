package handlers

import (
  "time"
  "net/http"

	"github.com/labstack/echo"
  "github.com/dgrijalva/jwt-go"

  "github.com/hiyali/katip-be/config"
)

// return errors.New("failed to connect database")

// Login with email & password
func UserLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

  db := config.GetDB()
  defer db.Close()

  var user config.User
  db.Where("email = ?", email).First(&user)

	if password == user.Password {
		// Set custom claims
		claims := &config.JwtCustomClaims{
      user.ID,
			user.Name,
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

package handlers

import (
  "time"
  "net/http"

  "github.com/labstack/echo"
  "github.com/dgrijalva/jwt-go"
  "golang.org/x/crypto/bcrypt"

  "github.com/hiyali/katip-be/config"
)

// Login with email & password
func AuthLogin(c echo.Context) (err error) {
  loginParams := new(config.JsonLogin)
  if err = c.Bind(loginParams); err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  }
  if err = c.Validate(loginParams); err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  }

  db := config.GetDB()
  defer db.Close()

  var user config.User
  if err = db.Where("email = ?", loginParams.Email).First(&user).Error; err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  }

  if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginParams.Password)); err != nil {
    return echo.ErrUnauthorized
  }

  // Set custom claims
  claims := &config.JwtCustomClaims{
    user.ID,
    user.Email,
    jwt.StandardClaims{
      ExpiresAt: time.Now().Add(time.Hour * 24 * 5).Unix(),
    },
  }

  // Create token with claims
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  // Generate encoded token and send it as response.
  t, err := token.SignedString([]byte("katip_known_secret"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  }

  return c.JSON(http.StatusOK, echo.Map{
    "token": t,
    "userInfo": config.JsonUser{
      ID: user.ID,
      Name: user.Name,
      Email: user.Email,
    },
  })
}

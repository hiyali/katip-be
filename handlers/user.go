package handlers

import (
  "time"
  "net/http"

  "github.com/labstack/echo"
  "github.com/dgrijalva/jwt-go"
  "golang.org/x/crypto/bcrypt"

  "github.com/hiyali/katip-be/config"
  "github.com/hiyali/katip-be/utils"
)

func UserRegister(c echo.Context) (err error) {
  user := new(config.JsonUser)
  if err = c.Bind(user); err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  }
  if err = c.Validate(user); err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  }

  db := config.GetDB()
  defer db.Close()

  var count uint
  if err := db.Model(&config.User{}).Where("email = ?", user.Email).Count(&count).Error; err != nil {
    return c.JSON(http.StatusInternalServerError, echo.Map{
      "message": err,
    })
  }
  if count > 0 {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": "This email already exist!",
    })
  }

  newPassword := utils.GeneratePassword(12, utils.SourceTypes{All:true})
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
  if err != nil {
    return c.JSON(http.StatusInternalServerError, echo.Map{
      "message": err,
    })
  }

  userInfo := config.User{
    CreatedAt: time.Now(),

    Name: user.Name,
    Email: user.Email,
    Password: string(hashedPassword),
  }

  if err := db.Create(&userInfo).Error; err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  } else {
    if err := utils.SendRegisterEmail(user.Email, user.Name, newPassword); err != nil {
      return c.JSON(http.StatusInternalServerError, echo.Map{
        "message": err,
      })
    }

    return c.JSON(http.StatusOK, echo.Map{
      "userInfo": config.JsonUser{
        ID: userInfo.ID,
        Name: userInfo.Name,
        Email: userInfo.Email,
      },
    })
  }
}

func UserGetOne(c echo.Context) (err error) {
  loginUser := c.Get("user").(*jwt.Token)
  claims := loginUser.Claims.(*config.JwtCustomClaims)

  db := config.GetDB()
  defer db.Close()

  var user config.User
  if err := db.Where("id = ?", claims.ID).First(&user).Error; err != nil {
    return c.JSON(http.StatusNotFound, echo.Map{
      "message": err,
    })
  } else {
    return c.JSON(http.StatusOK, config.JsonUser{
      claims.ID,
      user.Name,
      claims.Email,
    })
  }
}

func UserUpdateOne(c echo.Context) (err error) {
  user := new(config.JsonUserPut)
  if err = c.Bind(user); err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  }
  if err = c.Validate(user); err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  }

  loginUser := c.Get("user").(*jwt.Token)
  claims := loginUser.Claims.(*config.JwtCustomClaims)

  db := config.GetDB()
  defer db.Close()

  if len(user.Password) > 0 {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
      return c.JSON(http.StatusInternalServerError, echo.Map{
        "message": err,
      })
    }
    user.Password = string(hashedPassword)
  }

  user.UpdatedAt = time.Now()
  var userModel config.User
  if err := db.Model(&userModel).Where("id = ?", claims.ID).Updates(user).Error; err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  } else {
    return c.JSON(http.StatusOK, echo.Map{})
  }
}

package handlers

import (
  "time"
  "net/http"

  "github.com/labstack/echo"
  "github.com/dgrijalva/jwt-go"
  "golang.org/x/crypto/bcrypt"
  "github.com/roylee0704/gron"

  "github.com/hiyali/katip-be/config"
  "github.com/hiyali/katip-be/utils"
)

type (
  EmailTokenStore struct {
    Name        string
    Email       string
    Password    string
    CreatedAt   time.Time
    ExpiredAt   time.Time
  }
  Reminder struct {}
)

var emailTokenStore map[string]EmailTokenStore

const (
  TokenLength = 40
)

func (r Reminder) Run() {
  for token, val := range emailTokenStore {
    if len(val.Email) == 0 || val.ExpiredAt.Unix() < time.Now().Unix() {
      delete(emailTokenStore, token)
    }
  }
}

func init() {
  emailTokenStore = make(map[string]EmailTokenStore)

  // cron job
  r := Reminder{}
  c := gron.New()
  c.Add(gron.Every(2 * time.Minute), r)
  c.Start()
}

func UserRegister(c echo.Context) (err error) {
  user := new(config.JsonUserPost)
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

  token := utils.GenerateRandomStr(TokenLength, utils.SourceTypes{LowerLetters: true, UpperLetters: true, Digits: true})
  if err := utils.SendRegisterConfirmEmail(user.Email, user.Name, token); err != nil {
    return c.JSON(http.StatusInternalServerError, echo.Map{
      "message": err,
    })
  }

  emailTokenStore[token] = EmailTokenStore{
    Name: user.Name,
    Email: user.Email,
    Password: user.Password,
    CreatedAt: time.Now(),
    ExpiredAt: time.Now().Add(time.Duration(30) * time.Minute),
  }

  return c.JSON(http.StatusOK, echo.Map{
    "message": "Confirm email send!",
  })
}

func UserRegisterConfirm(c echo.Context) (err error) {
  Token := c.QueryParam("token")

  ets := emailTokenStore[Token]
  if len(ets.Email) == 0 || ets.ExpiredAt.Unix() < time.Now().Unix() {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": "Token expired!",
    })
  }

  db := config.GetDB()
  defer db.Close()

  var count uint
  if err := db.Model(&config.User{}).Where("email = ?", ets.Email).Count(&count).Error; err != nil {
    return c.JSON(http.StatusInternalServerError, echo.Map{
      "message": err,
    })
  }
  if count > 0 {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": "This email already exist!",
    })
  }

  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ets.Password), bcrypt.DefaultCost)
  if err != nil {
    return c.JSON(http.StatusInternalServerError, echo.Map{
      "message": err,
    })
  }

  userInfo := config.User{
    CreatedAt: time.Now(),

    Name: ets.Name,
    Email: ets.Email,
    Password: string(hashedPassword),
  }

  if err := db.Create(&userInfo).Error; err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
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

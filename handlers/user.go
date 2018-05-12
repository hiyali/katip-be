package handlers

import (
  "time"
  "net/http"
  "net/mail"
  "net/smtp"

  "github.com/labstack/echo"
  "github.com/dgrijalva/jwt-go"
  "golang.org/x/crypto/bcrypt"
  "github.com/scorredoira/email"
  "github.com/matcornic/hermes"

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

  newPassword := utils.GeneratePassword(12)
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

  db := config.GetDB()
  defer db.Close()

  if err := db.Create(&userInfo).Error; err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  } else {
    hermesEmail := hermes.Email{
      Body: hermes.Body{
        Name: user.Name,
        Intros: []string{
          "You have received this email because a password reset request for Hermes account was received.",
        },
        Actions: []hermes.Action{
          {
            Instructions: "Click the button below to reset your password:",
            Button: hermes.Button{
              Color: "#DC4D2F",
              Text:  "Reset your password: " + newPassword,
              Link:  "https://katip.hiyali.org/reset-password?token=d9729feb74992cc3482b350163a1a010",
            },
          },
        },
        Outros: []string{
          "If you did not request a password reset, no further action is required on your part.",
        },
        Signature: "Thanks",
      },
    }

    h := hermes.Hermes{
      Product: hermes.Product{
        Name: "Katip",
        Link: "https://katip.hiyali.org/",
        // Logo: "http://katip.hiyali.org/assets/logo.png",
        Copyright: "Copyright © 2018 Katip. All rights reserved.",
      },
    }
    emailBody, err := h.GenerateHTML(hermesEmail)
    if err != nil {
      return c.JSON(http.StatusInternalServerError, echo.Map{
        "message": err,
      })
    }

    m := email.NewHTMLMessage("Congratulations! You registered the Katip.", emailBody)
    m.From = mail.Address{Name: "From", Address: "katip@hiyali.org"}
    m.To = []string{"register-user@163.com"}
    /*
    if err := m.Attach("email.go"); err != nil {
      log.Fatal(err)
    }
    */

    // send it
    auth := smtp.PlainAuth("", "katip@hiyali.org", "non-secure", "smtp.hiyali.org")
    if err := email.Send("smtp.hiyali.org:25", auth, m); err != nil {
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
      "message": "Your password already send to your email address.",
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

  var userModel config.User
  if err := db.Model(&userModel).Where("id = ?", claims.ID).Updates(user).Error; err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  } else {
    return c.JSON(http.StatusOK, echo.Map{})
  }
}

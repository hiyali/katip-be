package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/roylee0704/gron"
	"golang.org/x/crypto/bcrypt"

	"github.com/hiyali/katip-be/config"
	"github.com/hiyali/katip-be/utils"
)

type (
	RegisterTokenStore struct {
		Name      string
		Email     string
		Password  string
		CreatedAt time.Time
		ExpiredAt time.Time
	}
	ResetTokenStore struct {
		Email     string
		CreatedAt time.Time
		ExpiredAt time.Time
	}
	Reminder struct{}
)

var registerTokenStore map[string]RegisterTokenStore
var resetTokenStore map[string]ResetTokenStore

const (
	TokenLength = 40
)

func (r Reminder) Run() {
	for token, val := range registerTokenStore {
		if len(val.Email) == 0 || val.ExpiredAt.Unix() < time.Now().Unix() {
			delete(registerTokenStore, token)
		}
	}

	for token, val := range resetTokenStore {
		if len(val.Email) == 0 || val.ExpiredAt.Unix() < time.Now().Unix() {
			delete(resetTokenStore, token)
		}
	}
}

func init() {
	registerTokenStore = make(map[string]RegisterTokenStore)
	resetTokenStore = make(map[string]ResetTokenStore)

	// cron job
	r := Reminder{}
	c := gron.New()
	c.Add(gron.Every(2*time.Minute), r)
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

	for _, val := range registerTokenStore {
		if val.Email == user.Email && val.ExpiredAt.Unix() > time.Now().Unix() {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "A register request being in the process with this email!",
			})
		}
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

	registerTokenStore[token] = RegisterTokenStore{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(time.Duration(30) * time.Minute),
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Confirm email send!",
	})
}

func UserRegisterConfirm(c echo.Context) (err error) {
	Token := c.QueryParam("token")

	registerTsItem := registerTokenStore[Token]
	if len(registerTsItem.Email) == 0 || registerTsItem.ExpiredAt.Unix() < time.Now().Unix() {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Token expired!",
		})
	}

	db := config.GetDB()
	defer db.Close()

	var count uint
	if err := db.Model(&config.User{}).Where("email = ?", registerTsItem.Email).Count(&count).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err,
		})
	}
	if count > 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "This email already exist!",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerTsItem.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err,
		})
	}

	userInfo := config.User{
		CreatedAt: time.Now(),

		Name:     registerTsItem.Name,
		Email:    registerTsItem.Email,
		Password: string(hashedPassword),
	}

	if err := db.Create(&userInfo).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Register confirmed",
		"userInfo": config.JsonUser{
			ID:    userInfo.ID,
			Name:  userInfo.Name,
			Email: userInfo.Email,
		},
	})
}

func UserGetInfo(c echo.Context) (err error) {
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
			user.AvatarUrl,
		})
	}
}

func UserUpdateInfo(c echo.Context) (err error) {
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

	var userModel config.User
	if err := db.Model(&userModel).Where("id = ?", claims.ID).Updates(user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err,
		})
	} else {
		return c.JSON(http.StatusOK, echo.Map{})
	}
}

func UserChangePassword(c echo.Context) (err error) {
	passwordInfo := new(config.JsonUserChangePassword)
	if err = c.Bind(passwordInfo); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err,
		})
	}
	if err = c.Validate(passwordInfo); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err,
		})
	}

	loginUser := c.Get("user").(*jwt.Token)
	claims := loginUser.Claims.(*config.JwtCustomClaims)

	db := config.GetDB()
	defer db.Close()

	var user config.User
	if err = db.Where("id = ?", claims.ID).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err,
		})
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordInfo.Password)); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Password is not correct.",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordInfo.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err,
		})
	}

	if err := db.Model(&user).Where("id = ?", claims.ID).Update("password", string(hashedPassword)).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err,
		})
	} else {
		return c.JSON(http.StatusOK, echo.Map{})
	}
}

func UserForgetPassword(c echo.Context) (err error) {
	emailInfo := new(config.JsonUserPost)
	if err = c.Bind(emailInfo); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err,
		})
	}
	if err = c.Validate(emailInfo); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err,
		})
	}

	for _, val := range resetTokenStore {
		if val.Email == emailInfo.Email && val.ExpiredAt.Unix() > time.Now().Unix() {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "A reset password request being in the process with this email!",
			})
		}
	}

	db := config.GetDB()
	defer db.Close()

	var user config.User
	if err = db.Where("email = ?", emailInfo.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err,
		})
	}

	token := utils.GenerateRandomStr(TokenLength, utils.SourceTypes{LowerLetters: true, UpperLetters: true, Digits: true})
	if err := utils.SendResetPasswordEmail(user.Email, user.Name, token); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err,
		})
	}

	resetTokenStore[token] = ResetTokenStore{
		Email:     emailInfo.Email,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(time.Duration(30) * time.Minute),
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

func UserResetPassword(c echo.Context) (err error) {
	resetInfo := new(config.JsonUserResetPassword)
	if err = c.Bind(resetInfo); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err,
		})
	}
	if err = c.Validate(resetInfo); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err,
		})
	}

	resetTsItem := resetTokenStore[resetInfo.Token]
	if len(resetTsItem.Email) == 0 || resetTsItem.ExpiredAt.Unix() < time.Now().Unix() {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Token expired!",
		})
	}

	db := config.GetDB()
	defer db.Close()

	var user config.User
	if err = db.Where("email = ?", resetTsItem.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err,
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(resetInfo.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err,
		})
	}

	if err := db.Model(&user).Where("email = ?", resetTsItem.Email).Update("password", string(hashedPassword)).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err,
		})
	} else {
		return c.JSON(http.StatusOK, echo.Map{})
	}
}

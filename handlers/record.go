package handlers

import (
  "net/http"

  "github.com/labstack/echo"
  "github.com/dgrijalva/jwt-go"

  "github.com/hiyali/katip-be/config"
)

// Authorization: Bearer {TOKEN_HERE}
// return errors.New("failed to connect database")

/*
@page  positive integer
@limit positive integer

return [record...]
*/

func RecordGetAllPageable(c echo.Context) (err error) {
  pageable := new(config.ParamPageable)
  if err = c.Bind(pageable); err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  }
  if err = c.Validate(pageable); err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  }

  user := c.Get("user").(*jwt.Token)
  claims := user.Claims.(*config.JwtCustomClaims)

  db := config.GetDB()
  defer db.Close()

  var records []config.Record
  db.Where("creator_id = ?", claims.ID).Limit(pageable.Limit).Offset(pageable.Limit * (pageable.Page - 1)).Find(&records)

  return c.JSON(http.StatusOK, records)
}

func RecordCreateOne(c echo.Context) (err error) {
  record := new(config.Record)
  if err = c.Bind(record); err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  }
  if err = c.Validate(record); err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  }

  user := c.Get("user").(*jwt.Token)
  claims := user.Claims.(*config.JwtCustomClaims)
  record.CreatorId = claims.ID

  db := config.GetDB()
  defer db.Close()

  if err := db.Create(&record).Error; err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  } else {
    return c.JSON(http.StatusOK, record)
  }
}

func RecordGetOne(c echo.Context) (err error) {
  id := c.Param("id")

  user := c.Get("user").(*jwt.Token)
  claims := user.Claims.(*config.JwtCustomClaims)

  db := config.GetDB()
  defer db.Close()

  var record config.Record
  if err := db.Where("creator_id = ? AND id = ?", claims.ID, id).First(&record).Error; err != nil {
    return c.JSON(http.StatusNotFound, echo.Map{
      "message": err,
    })
  } else {
    return c.JSON(http.StatusOK, record)
  }
}

func RecordUpdateOne(c echo.Context) (err error) {
  id := c.Param("id")

  record := new(config.Record)
  if err = c.Bind(record); err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  }
  if err = c.Validate(record); err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  }

  user := c.Get("user").(*jwt.Token)
  claims := user.Claims.(*config.JwtCustomClaims)

  db := config.GetDB()
  defer db.Close()

  var recordModel config.Record
  if err := db.Model(&recordModel).Where("creator_id = ? AND id = ?", claims.ID, id).Updates(record).Error; err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  } else {
    return c.JSON(http.StatusOK, echo.Map{})
  }
}

func RecordDeleteOne(c echo.Context) (err error) {
  id := c.Param("id")

  user := c.Get("user").(*jwt.Token)
  claims := user.Claims.(*config.JwtCustomClaims)

  db := config.GetDB()
  defer db.Close()

  var record config.Record
  if err := db.Where("creator_id = ? AND id = ?", claims.ID, id).Delete(&record).Error; err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  } else {
    return c.JSON(http.StatusOK, echo.Map{})
  }
}

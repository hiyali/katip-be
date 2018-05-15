package handlers

import (
  "time"
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
  db.Where("creator_id = ?", claims.ID).Limit(pageable.Limit).Offset(pageable.Skip).Find(&records)

  return c.JSON(http.StatusOK, records)
}

func RecordCreateOne(c echo.Context) (err error) {
  record := new(config.JsonRecord)
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

  recordInfo := config.Record{
    CreatorId: claims.ID,
    CreatedAt: time.Now(),

    Title: record.Title,
    IconUrl: record.IconUrl,
    Content: record.Content,
    Type: record.Type,
  }

  db := config.GetDB()
  defer db.Close()

  if err := db.Create(&recordInfo).Error; err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  } else {
    record.ID = recordInfo.ID
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

  record := new(config.JsonRecord)
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

  recordInfo := config.Record{
    Title: record.Title,
    IconUrl: record.IconUrl,
    Content: record.Content,
    Type: record.Type,
  }
  if err := db.Model(&recordInfo).Where("creator_id = ? AND id = ?", claims.ID, id).Updates(recordInfo).Error; err != nil {
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
